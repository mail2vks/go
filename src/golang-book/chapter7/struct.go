package main

import "fmt"
import "math"

type Circle struct {
    x, y, r float64
}

func area(c Circle) float64 {
    return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
    return math.Pi * c.r * c.r
}

func main() {
    c := Circle{0, 0, 3}
    fmt.Println(c.x, c.y, c.r)
    fmt.Println("Area is ", area(c))
    fmt.Println("Area is ", c.area())
}
