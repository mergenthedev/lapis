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
	switch in.(type) {
	case []float64:
		out, ok := in.([]float64)

		if !ok {
			fmt.Println(PrefixWarn)
		}

		return Vec3{out[0], out[1], out[2]}
	case Vec2:
		out, ok := in.(Vec2)

		if !ok {
			fmt.Println(PrefixWarn)
		}

		return Vec3{X: out.X, Y: out.Y, Z: 0}

	default:
		fmt.Println(PrefixWarn + "Couldn't convert to Vec3!")
		return Vec3{}
	}
}
