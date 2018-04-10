package util

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
)

var configCache map[string]*JsonConfig

type Config interface {
	String(path string) string
	Int64(path string) int64
	Interface(path string) interface{}
}

type JsonConfig struct {
	json string
}

func (Self *JsonConfig)Load(jsonFile string) error {
	data, err := ioutil.ReadFile(jsonFile)
	if err == nil {
		Self.json = string(data)
	}
	return err
}

func (Self *JsonConfig)String(path string) string {
	value := gjson.Get(Self.json, path)
	if value.Exists() {
		return value.String()
	}
	return ""
}

func (Self *JsonConfig)Int64(path string) int64 {
	value := gjson.Get(Self.json, path)
	if value.Exists() {
		return value.Int()
	}
	return 0
}

func (Self *JsonConfig)Interface(path string) interface{} {
	value := gjson.Get(Self.json, path)
	if value.Exists() {
		return value.Value()
	}
	return nil
}

func Configs(file string) Config {
	if configCache == nil {
		configCache = make(map[string]*JsonConfig)
	}

	if configCache[file] == nil{
		config := &JsonConfig{}
		err := config.Load(file)
		if err == nil {
			configCache[file] = config
		} else {
			panic(err)
		}
	}
	return configCache[file]
}