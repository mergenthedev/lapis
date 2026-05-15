package core

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	Pos        Vec3
	Near       float32
	Far        float32
	Fov        float32
	projection mgl32.Mat4
	pUniform   int32
	cam        mgl32.Mat4
	camUniform int32
}

func CreateCamera(fov float32, near float32, far float32, pos Vec3) Camera {
	var cam Camera = Camera{
		Pos:  Vec3{pos.X, pos.Y, pos.Z},
		Near: near,
		Far:  far,
		Fov:  fov,
	}

	cam.projection = mgl32.Perspective(mgl32.DegToRad(fov), 1, near, far)
	cam.pUniform = gl.GetUniformLocation(DefaultShaderPr, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(cam.pUniform, 1, false, &cam.projection[0])

	cam.cam = mgl32.LookAtV(mgl32.Vec3{pos.X, pos.Y, pos.Z}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	cam.camUniform = gl.GetUniformLocation(DefaultShaderPr, gl.Str("view\x00"))
	gl.UniformMatrix4fv(cam.camUniform, 1, false, &cam.cam[0])

	model := mgl32.Ident4()
	modelUniform := gl.GetUniformLocation(DefaultShaderPr, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	return cam
}

func (cam *Camera) UpdateCamera() {
	cam.cam = mgl32.LookAtV(mgl32.Vec3{cam.Pos.X, cam.Pos.Y, cam.Pos.Z}, mgl32.Vec3{0, 0, -1}.Add(mgl32.Vec3{cam.Pos.X, cam.Pos.Y, cam.Pos.Z}), mgl32.Vec3{0, 1, 0})
	gl.UniformMatrix4fv(cam.camUniform, 1, false, &cam.cam[0])
}
