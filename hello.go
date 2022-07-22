package main

import (
    "fmt"
    "strconv"
)

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
    
    for i := 1; i <= 10; i++ {
        
        if i % 2 == 0 {
            // 文字列に直す必要ある
            fmt.Println(strconv.Itoa(i) + "は偶数")
        } else {
            fmt.Println(strconv.Itoa(i) + "は奇数")
        }
    }
}