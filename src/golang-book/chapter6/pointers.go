package main

import "fmt"

func update(x *int) {
    fmt.Println(x)
    *x = 42
}

func square(x *int) {
    *x = *x * *x
}

func swap(y, z *int) {
    temp := *y
    *y = *z
    *z = temp
}

func main() {
    x := 32
    fmt.Println(&x)
    update(&x)
    fmt.Println(x)

    square(&x)
    fmt.Println(x)

    y := 3
    z := 2
    fmt.Println(y, z)
    swap(&y, &z)
    fmt.Println(y, z)
}
