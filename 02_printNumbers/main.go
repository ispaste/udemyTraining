package main

import "fmt"

func main()  {
    fmt.Printf("%d - %b - %x - %o \n", 42, 42, 42, 42)
    fmt.Printf("%d \t %b \t %X \n", 42, 42, 42)
    fmt.Printf("%d \t %b \t %#x \n", 42, 42, 42)
    fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
}