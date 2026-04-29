package core

import (
	"fmt"
	"log"
	"runtime"
	"time"

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

	glfw.SwapInterval(0)

	var GLVersionMajor, GLVersionMinor int32
	gl.GetIntegerv(gl.MAJOR_VERSION, &GLVersionMajor)
	gl.GetIntegerv(gl.MINOR_VERSION, &GLVersionMinor)

	fmt.Println(PrefixInfo + "OS: " + OS)
	fmt.Println(PrefixInfo+"CPU Threads:", runtime.NumCPU())
	fmt.Println(PrefixInfo + "GPU Vendor: " + gl.GoStr(gl.GetString(gl.VENDOR)))
	fmt.Printf(PrefixInfo+"OpenGL Version: %v.%v\n", GLVersionMajor, GLVersionMinor)

	gl.Viewport(0, 0, width, height)

	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	gl.DepthFunc(gl.LESS)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	defaultShader()

	return &Window{window}
}

var frame uint32 = 0
var lastCheckTime time.Time = time.Now()

func GetFPS() {
	frame++

	currentTime := time.Now()
	elapsed := currentTime.Sub(lastCheckTime).Seconds()

	if elapsed >= 1.0 {
		fmt.Printf("FPS: %d\n", frame)
		frame = 0
		lastCheckTime = currentTime
	}
}
