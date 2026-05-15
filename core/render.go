package core

import (
	"sync"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.4/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

var GenBuffers sync.Once
var vao, vbo, ebo uint32
var model mgl32.Mat4

func RenderLoop(window *Window, do func()) {
	for !window.ShouldClose() {
		glfw.PollEvents()
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		do()
		window.SwapBuffers()
	}
}

func DrawCube(obj map[string]any, tex uint32, deg float32, pos Vec3) {
	genCubeMeshBuffers()

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, tex)
	gl.Uniform1i(gl.GetUniformLocation(DefaultShaderPr, gl.Str("Texture\x00")), 0)

	model = mgl32.Translate3D(pos.X, pos.Y, pos.Z).Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(deg), mgl32.Vec3{0, 1, 0}))
	gl.UniformMatrix4fv(ModelLoc, 1, false, &model[0])

	gl.BindVertexArray(vao)
	gl.DrawElementsWithOffset(gl.TRIANGLES, 36, gl.UNSIGNED_INT, 0)
	gl.BindVertexArray(0)
}

func genCubeMeshBuffers() {
	GenBuffers.Do(func() {
		var vertices = []float32{
			// front
			-0.5, -0.5, 0.5, 0.0, 0.0,
			0.5, -0.5, 0.5, 1.0, 0.0,
			0.5, 0.5, 0.5, 1.0, 1.0,
			-0.5, 0.5, 0.5, 0.0, 1.0,

			// Back face (z = -0.5)
			0.5, -0.5, -0.5, 0.0, 0.0,
			-0.5, -0.5, -0.5, 1.0, 0.0,
			-0.5, 0.5, -0.5, 1.0, 1.0,
			0.5, 0.5, -0.5, 0.0, 1.0,

			// Top face (y = 0.5)
			-0.5, 0.5, -0.5, 0.0, 0.0,
			0.5, 0.5, -0.5, 1.0, 0.0,
			0.5, 0.5, 0.5, 1.0, 1.0,
			-0.5, 0.5, 0.5, 0.0, 1.0,

			// Bottom face (y = -0.5)
			-0.5, -0.5, 0.5, 0.0, 0.0,
			0.5, -0.5, 0.5, 1.0, 0.0,
			0.5, -0.5, -0.5, 1.0, 1.0,
			-0.5, -0.5, -0.5, 0.0, 1.0,

			// Right face (x = 0.5)
			0.5, -0.5, -0.5, 1.0, 0.0,
			0.5, 0.5, -0.5, 1.0, 1.0,
			0.5, 0.5, 0.5, 0.0, 1.0,
			0.5, -0.5, 0.5, 0.0, 0.0,

			// Left face (x = -0.5)
			-0.5, -0.5, 0.5, 1.0, 0.0,
			-0.5, 0.5, 0.5, 1.0, 1.0,
			-0.5, 0.5, -0.5, 0.0, 1.0,
			-0.5, -0.5, -0.5, 0.0, 0.0,
		}

		var indices = []uint32{
			0, 1, 2, 2, 3, 0, // Front
			4, 5, 6, 6, 7, 4, // Back
			8, 9, 10, 10, 11, 8, // Top
			12, 13, 14, 14, 15, 12, // Bottom
			16, 17, 18, 18, 19, 16, // Right
			20, 21, 22, 22, 23, 20, // Left
		}

		gl.GenVertexArrays(1, &vao)
		gl.BindVertexArray(vao)

		gl.GenBuffers(1, &vbo)
		gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
		gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

		gl.GenBuffers(1, &ebo)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

		gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 5*4, 0)
		gl.EnableVertexAttribArray(0)

		gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*4, 3*4)
		gl.EnableVertexAttribArray(1)

	})
}
