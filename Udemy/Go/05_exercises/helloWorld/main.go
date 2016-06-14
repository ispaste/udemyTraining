package main

import "fmt"

func main() {
	var name string
	fmt.Print("Please, write your name: ")
	fmt.Scanln(name)
	fmt.Println("Hello", name)
}
