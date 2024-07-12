package main

import (
	"log"
	"runtime"
	"runtime/debug"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	defer func() {
		if err := recover(); err != nil {
			log.Printf("go-main panic:%v stack:%v", err, string(debug.Stack()))
		}
	}()

	log.Printf("~~~~~~~~~ go-opengl: hello ~~~~~~~~~")

	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(1200, 800, "hellogl", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	for !window.ShouldClose() {
		//TODO, Do OpenGL stuff
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
