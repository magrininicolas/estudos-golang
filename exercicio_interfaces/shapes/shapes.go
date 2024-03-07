package shapes

import "math"

type Shape interface {
	Area() float64
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return math.Pow(s.Side, 2.0)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2.0)
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2.0
}

type EqtTriangle struct {
	Triangle
}

func (eT EqtTriangle) Area() float64 {
	return math.Pow(eT.Height, 2) * math.Sqrt(3) / 4.0
}
