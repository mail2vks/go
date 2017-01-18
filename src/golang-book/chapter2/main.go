package main

import "fmt"

func main() {
    fmt.Println("1 + 1 = ", 1+1)
    fmt.Println("1 + 1 = ", 1.0+1.0)
    fmt.Println(len("Hello, World"))
    fmt.Println("Hello, World"[1])
    fmt.Println("Hello, " + "World")

    var x int = 32
    y := 45
    var z = 97
    fmt.Println("x, y, z", x, y, z)

    fmt.Println("Now enter a number")
    var input float64
    fmt.Scanf("%f", &input)
    fmt.Println("Double of input is ", 2*input)
}
