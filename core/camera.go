package core

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	Pos        Vec3
	Near       uint32
	Far        uint32
	Fov        float32
	projection mgl32.Mat4
	pUniform   int32
	cam        mgl32.Mat4
	camUniform int32
}

func CreateCamera(fov float32, near float32, far float32, pos Vec3) {
	gl.UseProgram(DefaultShaderPr)

	projection := mgl32.Perspective(mgl32.DegToRad(fov), 1, near, far)
	projectionUniform := gl.GetUniformLocation(DefaultShaderPr, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	camera := mgl32.LookAtV(mgl32.Vec3{-4, 2, -3}, mgl32.Vec3{float32(pos.X), float32(pos.Y), float32(pos.Z)}, mgl32.Vec3{0, 1, 0})
	cameraUniform := gl.GetUniformLocation(DefaultShaderPr, gl.Str("view\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

	model := mgl32.Ident4()
	modelUniform := gl.GetUniformLocation(DefaultShaderPr, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])
}
