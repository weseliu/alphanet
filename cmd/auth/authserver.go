package main

import (
	"encoding/json"
	"fmt"
	"github.com/weseliu/alphanet/cmd/global/ado"
	"github.com/weseliu/alphanet/db"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/weseliu/alphanet/cmd/global/interfaces"
	"github.com/weseliu/alphanet/util"
)

type authRequest struct {
	UserChannel string `json:"user_channel"`
	Acc         string `json:"acc"`
	Pwd         string `json:"pwd"`
	AccessToken string `json:"access_token"`
}

type authResponse struct {
	RetCode       string `json:"ret_code"`
	RetMsg        string `json:"ret_msg"`
	IdentityToken string `json:"identity_token"`
	ServerUrl     string `json:"server_url"`
}

type AuthResult int

const (
	AuthResultSuccess AuthResult = iota
	AuthResultParamsError
	AuthResultRegisterFail
	AuthResultAccountError
	AuthResultTokenError
)

func (Self AuthResult) String() string {
	return strconv.Itoa((int)(Self))
}

func (Self AuthResult) Description() string {
	switch Self {
	case AuthResultSuccess:
		return "AuthResultSuccess"
	case AuthResultParamsError:
		return "AuthResultParamsError"
	case AuthResultRegisterFail:
		return "AuthResultRegisterFail"
	case AuthResultAccountError:
		return "AuthResultAccountError"
	case AuthResultTokenError:
		return "AuthResultTokenError"
	}
	return "Error"
}

var config util.Config

func sendAuthResult(w http.ResponseWriter, code AuthResult, token string) {
	var rsp = authResponse{
		RetCode:       code.String(),
		RetMsg:        code.Description(),
		IdentityToken: token,
		ServerUrl:     config.String("connect_server"),
	}

	data, err := json.Marshal(rsp)
	if err == nil {
		io.WriteString(w, string(data))
	}
}

func authServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	req.ParseForm()

	body, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	fmt.Printf("%s\n", body)

	w.WriteHeader(200)

	authReq := &authRequest{}
	err := json.Unmarshal(body, authReq)
	if err != nil || len(authReq.Acc) == 0 || len(authReq.UserChannel) == 0 {
		sendAuthResult(w, AuthResultParamsError, "")
		return
	}

	regInfo := ado.Register().GetUser(authReq.Acc, authReq.UserChannel)
	if regInfo == nil {
		regInfo = &ado.RegisterModel{
			Account:  authReq.Acc,
			Channel:  authReq.UserChannel,
			Password: authReq.Pwd,
		}

		if success := ado.Register().AddUser(regInfo); success != true {
			sendAuthResult(w, AuthResultRegisterFail, "")
			return
		}
	}

	if regInfo.Password != authReq.Pwd {
		sendAuthResult(w, AuthResultAccountError, "")
		return
	}

	token := &interfaces.IdentityToken{
		Account:regInfo.Account,
		Channel:regInfo.Channel,
		Time:time.Now(),
	}

	identityToken, err := token.Encrypt()
	if err != nil {
		sendAuthResult(w, AuthResultTokenError, "")
		return
	}
	sendAuthResult(w, AuthResultSuccess, identityToken)
}

func main() {
	config = util.Configs("./conf/auth.json")
	db.Instance().Open(config.String("dsn"))

	http.HandleFunc("/auth", authServer)
	err := http.ListenAndServe(config.String("port"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
