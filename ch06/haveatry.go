package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	*Point
	Color color.RGBA
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{1, 2}
	q := Point{3, 4}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))

	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	k := Point{1, 2}
	(&k).ScaleBy(3)
	fmt.Println(k)

	//编译器隐式地获取变量的地址
	p.ScaleBy(4)
	fmt.Println(p)

	red := color.RGBA{R: 255, A: 255}
	blue := color.RGBA{B: 255, A: 255}
	p1 := ColoredPoint{&Point{1, 1}, red}
	q1 := ColoredPoint{&Point{5, 4}, blue}
	fmt.Println(p1.Distance(*q1.Point))
	q1.Point = p1.Point
	//q1.ScaleBy(2)
	p1.ScaleBy(2)
	fmt.Println(*p1.Point, *q1.Point)

}
