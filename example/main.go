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
	fmt.Println(core.Get(scene, "cam.near"))

	texture := core.LoadImage("jackashi.jpg", core.LINEAR)

	gl.ClearColor(0, 0, 0, 0)

	core.CreateCamera(45, 0.1, 1000, core.Vec3{0, 0, 0})

	fmt.Println(scene.Script)
	fmt.Println(scene.Name)

	core.InitGUI()

	for !window.ShouldClose() {
		glfw.PollEvents()
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		core.DrawCube(scene.Objects, texture)
		window.SwapBuffers()

		core.GetFPS()
	}

	core.End()
}
