package main

type JsonCommand struct {
	name  string
	param string
	data  string
}

type UserAuth struct {
	name     string
	password string
	channel  string
	deviceId string
	platform string
}
