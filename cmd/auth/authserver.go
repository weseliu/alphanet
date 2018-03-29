package main

import (
	"net/http"
	"log"
	"io"
	"io/ioutil"
	"fmt"
)

func authServer(w http.ResponseWriter, req *http.Request)  {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	req.ParseForm()

	result, _:= ioutil.ReadAll(req.Body)
	req.Body.Close()
	fmt.Printf("%s\n", result)

	w.WriteHeader(200)
	io.WriteString(w, string(result[:]))
}

func main() {
	http.HandleFunc("/auth", authServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}