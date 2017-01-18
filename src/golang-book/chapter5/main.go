package main

import "fmt"

func main() {

    var x [5]float64
    x[0] = 1
    x[1] = 2
    x[2] = 3
    x[3] = 4
    x[4] = 5

    var total float64 = 0
    for i := 0; i < len(x); i++ {
        total += x[i]
    }

    fmt.Println("Average of numbers is ", total/float64(len(x)))

    // Improving for loop
    total = 0
    for _, value := range x {
        total += value
    }

    fmt.Println("Average of numbers is ", total/float64(len(x)))

    // Create a slice out of x
    fmt.Println(x[:])
    fmt.Println(x[:4])

    //declare map
    m := make(map[string]int)
    m["India"] = 1
    fmt.Println(m)

    if v, ok := m["India"]; ok {
        fmt.Println("Value in map against key is ", v)
    }

}
