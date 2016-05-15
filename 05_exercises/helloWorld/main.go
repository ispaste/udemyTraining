package main

import "fmt"

func main() {
	fmt.Print("Please, write your name: ")
	fmt.Scanln(name, error)
	fmt.Fprintln("Hello " + name)
}
