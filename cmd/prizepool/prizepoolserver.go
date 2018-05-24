package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/weseliu/alphanet/util"
	"io"
	"time"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type prizeResponse struct {
	Bullet       string `json:"bullet"`
	Times[]        string `json:"times"`
	Values[] 	float32 `json:"values"`
}

var config util.Config

func sendResult(w http.ResponseWriter, bullet string, times[] string, values[] float32) {
	var rsp = prizeResponse{
		Bullet:       bullet,
		Times:        times,
		Values: values,
	}

	data, err := json.Marshal(rsp)
	if err == nil {
		io.WriteString(w, string(data))
	}
}

func getLogFiles(startTime string, endTime string) (files[] string) {
	files = make([]string, 0)
	now := time.Now()
	start, _ := time.Parse("2006-01-02 15:04:05", startTime)
	end, _ := time.Parse("2006-01-02 15:04:05", endTime)
	startDay := start.Day()
	endDay := end.Day()
	if start.Day() > end.Day(){
		startDay = end.Day()
		endDay = start.Day()
	}

	for i := startDay; i <= endDay; i++ {
		name := config.String("log_path")
		if i != now.Day() {
			name = config.String("log_path") + time.Date(start.Year(), start.Month(), i, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
		}
		print("getLogFiles:", name)
		files = append(files, name)
	}

	return files
}

func prizeServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	req.ParseForm()

	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	var bullet = req.Form.Get("bullet")
	var startTime = req.Form.Get("startTime")
	var endTime = req.Form.Get("endTime")

	w.WriteHeader(200)

	var times = make([]string, 0)
	var values = make([]float32, 0)

	if len(bullet) == 0 {
		sendResult(w, bullet, times, values)
		return
	}

	var files = getLogFiles(startTime, endTime)
	for _, file := range files {
		fi, err := os.Open(file)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}

		br := bufio.NewReader(fi)
		for {
			a, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			line := string(a)
			items := strings.Split(line, "|")
			if len(items) == 6 {
				if items[0] >= startTime && items[0] <= endTime && items[2] == bullet{
					times = append(times, items[0])
					v, _ := strconv.ParseFloat(items[3], 32)
					values = append(values, float32(v))
				}
			}
		}
		fi.Close()
	}
	sendResult(w, bullet, times, values)
}

func main() {
	config = util.Configs("./conf/prizepool.json")
	http.HandleFunc("/prize", prizeServer)
	err := http.ListenAndServe(config.String("port"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
