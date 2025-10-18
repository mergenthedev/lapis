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

	fmt.Println(core.Get(scene, "obj2.energy"))
	core.Edit(scene, "obj2.energy", 5)
	fmt.Println(core.Get(scene, "obj2.energy"))

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		window.SwapBuffers()
		glfw.PollEvents()
	}

	core.End()
}
