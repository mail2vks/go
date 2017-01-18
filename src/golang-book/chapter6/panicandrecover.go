package main

import "fmt"

func main() {
    defer func() {
        rv := recover()
        fmt.Println("Got message during panic", rv)
    }()
    panic("I am panicking")
}
