package core

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.4/glfw"
)

type Window struct {
	*glfw.Window
}

var DefaultWindow *glfw.Window

func CreateWindow(width int32, height int32, title string, resizable int) *Window {
	glfw.WindowHint(glfw.Resizable, resizable)
	window, err := glfw.CreateWindow(int(width), int(height), title, nil, nil)
	DefaultWindow = window

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

	switch Engine.VSync {
	case true:
		glfw.SwapInterval(1)
	case false:
		glfw.SwapInterval(0)
	}

	switch Engine.MSAA_Enabled {
	case true:
		gl.Enable(gl.MULTISAMPLE)
	case false:
		gl.Disable(gl.MULTISAMPLE)
	}

	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	gl.DepthFunc(gl.LESS)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.ClearColor(0.0, 0.0, 0.0, 0.0)

	defaultShader()
	gl.UseProgram(DefaultShaderPr)

	return &Window{window}
}
