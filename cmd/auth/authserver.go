package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
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

func authServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	req.ParseForm()

	result, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	fmt.Printf("%s\n", result)

	w.WriteHeader(200)
	var rsp = authResponse{
		RetCode : "0",
		RetMsg: "",
		IdentityToken : "asfd233asdg33asdg",
		ServerUrl : "ws://127.0.0.1:8801/login",
	}
	data, _ := json.Marshal(rsp)
	io.WriteString(w, string(data))
}

func main() {
	http.HandleFunc("/auth", authServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
