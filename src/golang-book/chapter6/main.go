package main

import "fmt"

func main() {
    x := []float64{48, 96, 86, 68,
        57, 82, 63, 70,
        37, 34, 83, 27,
        19, 97, 9, 17,
    }
    fmt.Println("Average of numbers is ", avg(x))

    fmt.Println(add(1, 2, 3, 4, 5))
}

func avg(x []float64) float64 {
    var total float64 = 0
    for _, value := range x {
        total += value
    }
    return total / float64(len(x))
}

func add(input ...int) int {
    total := 0
    for _, value := range input {
        total += value
    }
    return total
}
