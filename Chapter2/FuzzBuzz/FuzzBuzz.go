package main

import "fmt"

func main(){
    var b bool
    for i := 1; i < 101; i++ {
        b = false
        if i%3 == 0{
            fmt.Println("Fizz")
            b = true
        }
        if i%5 == 0{
            fmt.Println("Buzz")
            b = true
        }
        if b == false{
            fmt.Println(i)
        }
        fmt.Println()
    }
}