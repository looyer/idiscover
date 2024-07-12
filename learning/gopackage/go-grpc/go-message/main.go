package main

import (
	"encoding/json"
	"log"
	"meproto"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Printf("~~~~~~~~~ go-message ~~~~~~~~~")

	errMessage := &meproto.ErrorMessage{
		Code:        meproto.ErrorCode_Error_Config,
		ErrorString: "Sorry! Can't find config!",
	}

	jsonstr, _ := json.Marshal(errMessage)
	log.Printf("errMessage json: %v", string(jsonstr)) //errMessage json: {"Code":2,"ErrorString":"Sorry! Can't find config!"}

	err2 := &meproto.ErrorMessage{}
	json.Unmarshal(jsonstr, err2)
	log.Printf("err2: %v", err2) //err2: Code:Error_Config ErrorString:"Sorry! Can't find config!"
}
