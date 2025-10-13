package core

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Window struct {
	*glfw.Window
}

func CreateWindow(width int32, height int32, title string, resizable int) *Window {
	glfw.WindowHint(glfw.Resizable, resizable)
	window, err := glfw.CreateWindow(int(width), int(height), title, nil, nil)

	if err != nil {
		log.Fatal(PrefixErr + "Can't create window!")
	}

	window.MakeContextCurrent()

	var GLVersionMajor, GLVersionMinor int32
	gl.GetIntegerv(gl.MAJOR_VERSION, &GLVersionMajor)
	gl.GetIntegerv(gl.MINOR_VERSION, &GLVersionMinor)

	fmt.Println(PrefixInfo + "OS: " + OS)
	fmt.Println(PrefixInfo+"CPU Threads:", runtime.NumCPU())
	fmt.Println(PrefixInfo + "GPU Vendor: " + gl.GoStr(gl.GetString(gl.VENDOR)))
	fmt.Printf(PrefixInfo+"OpenGL Version: %v.%v\n", GLVersionMajor, GLVersionMinor)

	gl.Viewport(0, 0, width, height)

	return &Window{window}
}
