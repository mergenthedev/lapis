package core

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const OS = runtime.GOOS

type Config struct {
	Debug bool
	VSync bool
}

var Engine Config

func init() {
	runtime.LockOSThread()
}

func Init(config Config) {
	if err := glfw.Init(); err != nil {
		log.Fatal(PrefixErr + "Cannot init GLFW!")
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.AlphaBits, 8)

	if err := gl.Init(); err != nil {
		log.Fatal(PrefixErr + "Cannot init OpenGL!")
	}

	Engine = config
}

func End() {
	glfw.Terminate()
}
