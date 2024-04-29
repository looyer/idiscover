package main

import (
	"log"

	"gocv.io/x/gocv"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)

	log.Printf("~~~~~~~~~ hello go-opencv ~~~~~~~~~~")
	log.Printf("gocv version: %s\n", gocv.Version())
	log.Printf("opencv lib version: %s\n", gocv.OpenCVVersion())
}
