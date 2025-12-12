package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type square struct {
	side float64
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perimeter() float64 {
	return 4 * s.side
}

func calculateArea(s shape) {
	area := s.area()
	fmt.Printf("Area of given shape is %v\n", area)
}

func calculatePerimeter(s shape) (float64, string) {
	c,ok := s.(circle)
	if !ok{
		fmt.Println("We have not passed a circle!")
	} else if c.radius>0 {
		fmt.Println("We have passed a circle!")
	}
	switch v := s.(type) {
	case square:
		return v.perimeter(), "passed"
	case circle:
		return 0.00, "Not implemented"
	}
	return 0.00, "Unknown type"
}

func main() {
	calculateArea(circle{
		radius: 10,
	})

	calculateArea(square{
		side: 10,
	})

	squarePerimeter, squareString := calculatePerimeter(square{
		side: 10,
	})

	fmt.Printf("The square is %v and string is %v\n", squarePerimeter, squareString)

	circlePerimeter, circleString := calculatePerimeter(circle{
		radius: 10,
	})

	fmt.Printf("The square is %v and string is %v\n", circlePerimeter, circleString)
}
