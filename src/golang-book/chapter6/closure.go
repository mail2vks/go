package main

import "fmt"

func evenGenerator() func() int {
    x := 0
    return func() int {
        x += 2
        return x
    }

}

func main() {
    f := evenGenerator()
    fmt.Println(f())
    fmt.Println(f())
    fmt.Println(f())
    fmt.Println(f())
    fmt.Println(f())
}
