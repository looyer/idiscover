package main

import (
	"encoding/json"
	"log"
	"meproto"

	"google.golang.org/protobuf/encoding/protojson"
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

	msg := &meproto.Building{
		UUID: 1520001,
		ID:   101,
		Factory: &meproto.Factory{
			ListOK:   []int32{1, 2, 5},
			ListIn:   []int32{3, 7, 7},
			DoneTime: 2004,
		},
	}
	jsonstr, _ = json.Marshal(msg)
	log.Printf("building json:%v", string(jsonstr))

	msg2 := &meproto.Building{}
	// json.Unmarshal(jsonstr, msg2)
	protojson.Unmarshal(jsonstr, msg2)

	factory := msg2.GetFactory()
	land := msg2.GetLand()

	log.Printf("msg2: %v factory:%v land:%v", msg2, factory, land)
}
