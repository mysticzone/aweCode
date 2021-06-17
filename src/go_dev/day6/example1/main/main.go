package main

import (
	"fmt"
)

type IPoint interface {
	X() float64
	Y() float64
	SetX(float642 float64)
	SetY(float642 float64)
}

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	// p := New(Point)
	p := new(Point)
	p.SetX(x)
	p.SetY(y)
	return p
}

func (p *Point) X() float64 {
	return p.x
}

func (p *Point) Y() float64 {
	return p.y
}

func (p *Point) SetX(x float64) {
	p.x = x
}

func (p *Point) SetY(y float64) {
	p.y = y
}

type Point3D struct {
	Point
	z float64
}

func NewPoint3D(x float64, y float64, z float64) *Point3D {
	p := new(Point3D)
	p.SetX(x)
	p.SetY(y)
	p.SetZ(z)
	return p
}

func (p *Point3D) Z() float64 {
	return p.z
}

func (p *Point3D) SetZ(z float64) {
	p.z = z
}

func main() {
	points := make([]IPoint, 0)

	p1 := NewPoint(3, 4)
	p2 := NewPoint3D(1, 2, 3)

	points = append(points, p1, p2)
	for _, p := range points {
		fmt.Println(fmt.Sprintf("(%.2f, %.2f)", p.GetX(), p.GetY))
	}
}
