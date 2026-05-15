package core

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
)

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

type Vec2 struct {
	X float32
	Y float32
}

func ToVec3(in any) Vec3 {
	switch val := in.(type) {
	case [3]float32:
		return Vec3{val[0], val[1], val[2]}
	case []float32:
		return Vec3{val[0], val[1], val[2]}
	case Vec2:
		return Vec3{X: val.X, Y: val.Y, Z: 0}
	case mgl32.Vec3:
		return Vec3{X: float32(val.X()), Y: float32(val.Y()), Z: float32(val.Z())}
	default:
		fmt.Println(PrefixWarn + "Couldn't convert to Vec3!")
		return Vec3{}
	}
}
