package main

import (
	"fmt"
	"math"
)

type rect2 struct {
	width, height float64
}

type circle2 struct {
	radius float64
}

func (r rect2) area2() float64 {
	return r.width * r.height
}

func (r rect2) perim2() float64 {
	return 2*r.width + 2*r.height
}

func (c circle2) area2() float64 {
	return math.Pi * c.radius * c.radius // πr^2
}

func (c circle2) perim2() float64 {
	return 2 * math.Pi * c.radius //2πr
}

func measureRect(g rect2) {
	fmt.Println(g)
	fmt.Println(g.area2())
	fmt.Println(g.perim2())
}

func measureCircle(g circle2) {
	fmt.Println(g)
	fmt.Println(g.area2())
	fmt.Println(g.perim2())
}

func without() {
	r := rect2{width: 3, height: 4}
	c := circle2{radius: 5}

	measureRect(r)
	measureCircle(c)

}
