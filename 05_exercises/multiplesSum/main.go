package main

import "fmt"

func main()  {
    sum := 0
    for i := 1; i < 1000; i++ {
        if i % 3 == 0 || i % 5 == 0 {
            sum = sum + i
        }
    }
    fmt.Println("The sum of all multiples of 3 and 5 below 1000 is", sum)
}