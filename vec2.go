package main

import (
	"math"
)

type vec2 struct {
	X, Y float64
}

func (v vec2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)

}

func (v vec2) Normalized() vec2 {
	xs := v.X / v.Length()
	ys := v.Y / v.Length()
	return vec2{xs, ys}

}

func (v vec2) Multiply(s float64) vec2 {
	xs := v.X * s
	ys := v.Y * s
	return vec2{xs, ys}
}

func (v vec2) Added(a vec2) vec2 {
	xa := v.X + a.X
	ya := v.Y + a.Y
	return vec2{xa, ya}
}

func (v vec2) DistanceTo(p vec2) float64 {
	return math.Sqrt((v.X-p.X)*(v.X-p.X) + (v.Y-p.Y)*(v.Y-p.Y))
}
