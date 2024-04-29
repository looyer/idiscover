package main

import (
	"log"

	"gocv.io/x/gocv"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)

	log.Printf("~~~ opencv tutorial01 showimage ~~~")

	mat := gocv.IMRead("./workpoint.jpg", 0)
	if mat.Empty() {
		return
	}

	window := gocv.NewWindow("Tutorial01-ShowImage")
	window.IMShow(mat)

	k := window.WaitKey(0)
	if k == int('q') {
		return
	}
}
