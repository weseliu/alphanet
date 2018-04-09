package util

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
)

var jsonCache map[string]string

func GetJsonValue(file string, path string) gjson.Result {
	if jsonCache == nil {
		jsonCache = make(map[string]string)
	}

	if jsonCache[file] == ""{
		data, err := ioutil.ReadFile("F:/GoProjects/src/github.com/weseliu/alphanet/cmd/auth/conf/" + file)
		if err != nil {
			panic(err)
		}
		jsonCache[file] = string(data)
	}

	return gjson.Get(jsonCache[file], path)
}