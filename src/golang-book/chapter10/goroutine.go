package main

import "fmt"
import "time"
import "math/rand"

func f() {
    for i := 0; i <= 10; i++ {
        fmt.Println(i)
        amt := time.Duration(rand.Intn(250))
        time.Sleep(time.Millisecond * amt)
    }
}

func main() {
    for j := 0; j <= 20; j++ {
        go f()
    }

    var input string
    fmt.Println("Please provide some input ")
    fmt.Scanln(&input)
    fmt.Println("Input is", input)
}
