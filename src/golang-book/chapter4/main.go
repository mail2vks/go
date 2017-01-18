package main

import "fmt"

func main() {
    i := 0
    for i <= 10 {
        fmt.Println(i)
        i += 1
    }

    for i := 0; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Println(i, " is even")
        } else {
            fmt.Println(i, " is odd")
        }
    }
}
