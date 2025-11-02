package core

import (
	"fmt"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

type Vec2 struct {
	X float64
	Y float64
}

func ToVec3(in any) Vec3 {
	switch val := in.(type) {
	case [3]float64:
		return Vec3{val[0], val[1], val[2]}
	case []float64:
		return Vec3{val[0], val[1], val[2]}
	case Vec2:
		return Vec3{X: val.X, Y: val.Y, Z: 0}
	default:
		fmt.Println(PrefixWarn + "Couldn't convert to Vec3!")
		return Vec3{}
	}
}
