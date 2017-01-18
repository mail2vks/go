package main

import "fmt"
import "math"

type Shape interface {
    area() float64
}

type Circle struct {
    r float64
}

type Square struct {
    s float64
}

func (c Circle) area() float64 {
    return math.Pi * c.r * c.r
}

func (s Square) area() float64 {
    return s.s * s.s
}

func getAllAreas(shapes ...Shape) float64 {
    var totalArea float64
    for _, s := range shapes {
        totalArea += s.area()
    }
    return totalArea
}

func main() {
    c := Circle{3}
    s := Square{4}
    fmt.Println("Total area is ", getAllAreas(&c, &s))
}
