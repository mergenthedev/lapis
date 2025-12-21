package main

import (
	"fmt"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/mergenthedev/lapis/core"
)

func main() {
	core.Init(core.Config{
		Debug: false,
		VSync: false,
	})

	var window *core.Window = core.CreateWindow(800, 800, "Lapis", core.FALSE)

	var scene = core.LoadScene("scene.toml")

	willBeAVec3 := []float64{9, 1, 5}
	thisToo := core.Vec2{X: 2, Y: 0}

	fmt.Println(core.ToVec3(willBeAVec3))
	fmt.Println(core.ToVec3(thisToo))

	fmt.Println(" ")

	fmt.Println(core.Get(scene.Cameras, "cam.near"))

	shader := core.CreateCustomShader("vertex.glsl", "fragment.glsl")

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		core.Draw(scene.Objects, shader)
		window.SwapBuffers()
		glfw.PollEvents()
	}

	core.End()
}
