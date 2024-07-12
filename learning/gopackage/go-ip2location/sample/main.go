package main

import (
	"log"
	"runtime/debug"

	"github.com/ip2location/ip2location-go/v9"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic:%v stack:%v", err, string(debug.Stack()))
		}
	}()

	log.Printf("~~~~~~~~~ go-ip2location: sample ~~~~~~~~~")

	db, err := ip2location.OpenDB("./IP2LOCATION-LITE-DB1.BIN")
	if err != nil {
		log.Fatalf("ip2location err:%v", err)
	}
	result, err := db.Get_all("185.244.208.95")
	if err != nil {
		log.Fatalf("ip2location Get_all:%v", err)
	}
	log.Printf("region:%v short:%v", result.Country_long, result.Country_short)
}
