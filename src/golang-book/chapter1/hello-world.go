package main

import "fmt"

func main() {
    numWritten, err := fmt.Println("")
    fmt.Println(numWritten)
    fmt.Println(err)
}
