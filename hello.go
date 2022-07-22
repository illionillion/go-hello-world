package main

import "fmt"

func main() {
    fmt.Println("Hello!");
    fmt.Println("World!");

    val1 := 0

    val1 = 12

    
    val1 += 10
    // val1 -= 10

    fmt.Println(val1)
    
    if val1 >= 10 {
        fmt.Println("val1 >= 10")
    } else {
        fmt.Println("val1 < 10")
    }
}