package main

import (
	"encoding/json"
	"fmt"
	"github.com/weseliu/alphanet/cmd/ado"
	"github.com/weseliu/alphanet/db"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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

func sendAuthResult(w http.ResponseWriter, code string, msg string, token string) {
	var rsp = authResponse{
		RetCode:       code,
		RetMsg:        msg,
		IdentityToken: token,
		ServerUrl:     "ws://127.0.0.1:8801/login",
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
	if err != nil || len(authReq.Acc) == 0 || len(authReq.UserChannel) == 0{
		sendAuthResult(w, "-1", "params format error!", "")
		return
	}

	regInfo := ado.Register().GetUser(authReq.Acc, authReq.UserChannel)
	if regInfo == nil {
		regInfo = &ado.RegisterModel{
			Account:  authReq.Acc,
			Channel:  authReq.UserChannel,
			Password: authReq.Pwd,
		}

		if success := ado.Register().AddUser(regInfo); success != true{
			sendAuthResult(w, "-2", "register user fail!", "")
			return
		}
	}

	sendAuthResult(w, "0", "", util.Md5(regInfo.Account + regInfo.Channel))
}

func main() {
	db.Instance().Open("root:@tcp(localhost:3306)/alphanet?charset=utf8")

	http.HandleFunc("/auth", authServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
