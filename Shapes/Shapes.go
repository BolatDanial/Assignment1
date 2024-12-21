package Shapes

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Length float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Length
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Length)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * (3.14 * c.Radius)
}

type Square struct {
	Length float64
}

func (r Square) Area() float64 {
	return r.Length * r.Length
}
func (r Square) Perimeter() float64 {
	return 4 * r.Length
}

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func (t Triangle) Area() float64 {
	p := (t.SideA + t.SideB + t.SideC) / 2
	return math.Sqrt(p * (p - t.SideA) * (p - t.SideB) * (p - t.SideC))
}
func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

func PrintShapeDetails(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}
