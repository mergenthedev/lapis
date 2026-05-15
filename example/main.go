package main

import (
	"fmt"

	"github.com/mergenthedev/lapis/core"
)

func main() {
	core.Init(core.Config{
		Debug:        false,
		VSync:        false,
		MSAA_Enabled: true,
		MSAA_Samples: 16,
	})

	var window *core.Window = core.CreateWindow(800, 800, "Lapis", core.FALSE)

	var scene = core.LoadScene("scene.toml")

	fmt.Println(scene.Objects)

	texture := core.LoadImage("mm.jpg", core.LINEAR)
	texture2 := core.LoadImage("jackashi.jpg", core.LINEAR)
	texture3 := core.LoadImage("stella.png", core.LINEAR)

	cam := core.CreateCamera(90, 0.1, 1000, core.Vec3{X: 0, Y: 0, Z: 2})

	var deg float32 = 0

	core.RenderLoop(window, func() {
		deg -= 0.005
		//cam.Pos.Y += 0.0001
		core.DrawCube(scene.Objects, texture, deg, core.Vec3{X: 1, Y: 1, Z: -0.5})
		core.DrawCube(scene.Objects, texture, deg, core.Vec3{X: -3, Y: 1.7, Z: -3})
		core.DrawCube(scene.Objects, texture2, deg, core.Vec3{X: -2, Y: 0, Z: -1.5})
		core.DrawCube(scene.Objects, texture2, deg, core.Vec3{X: -1.8, Y: -1.6, Z: -2})
		core.DrawCube(scene.Objects, texture3, deg, core.Vec3{X: .3, Y: -0.8, Z: 0})
		core.DrawCube(scene.Objects, texture3, deg, core.Vec3{X: -0.7, Y: 0.8, Z: -2.3})
		cam.UpdateCamera()
	})

	core.End()
}
