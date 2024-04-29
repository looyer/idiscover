package main

import (
	"errors"
	"fmt"
	"log"
	"runtime/debug"
	"unicode/utf8"
)

var ErrCold error = fmt.Errorf("can't boiling water ")
var ErrMilk error = fmt.Errorf("can't use milk")

func makeTea(c int32) error {
	if c == 2 {
		return fmt.Errorf("c err-2:%w", ErrCold) //wrap error is very useful. use errors.Is(e, target) check.
	}
	return nil
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic:%v stack:%v", err, string(debug.Stack()))
		}
	}()

	log.Printf("~~~~~~~~~ go-byfree ~~~~~~~~~")

	stru8 := "hello,你好，！emoji-😀ok."
	// for k, v := range stru8 {
	// 	log.Printf("start:%v rune:%v", k, v) //k:rune start index, v:rune
	// }
	log.Printf("stru8:%v size:%v runes:%v", stru8, len(stru8), utf8.RuneCountInString(stru8))

	if e := makeTea(2); e != nil {
		if errors.Is(e, ErrCold) {
			log.Printf("can't use cold water to make tea.")
		}
	} else {
		log.Printf("tea is maked.")
	}

	//decimal and integer change
	var v0 interface{} = int32(5)
	fa, ok := v0.(float32)
	log.Printf("v0:%v fa:%v ok:%v", v0, fa, ok)

	fb := v0.(float32)
	log.Printf("fb:%v", fb)
}
