package main

import "fmt"

func main() {
	for i := 1; i < 50; i++ {
		devBy3and5 := i%3 == 0 && i%5 == 0
		devBy3 := i%3 == 0
		devBy5 := i%5 == 0

		switch {
		case devBy3and5:
			fmt.Print("Fizz Buzz,")
		case devBy3:
			fmt.Print("Fizz,")
		case devBy5:
			fmt.Print("Buzz,")
		default:
			fmt.Printf("%d,", i)
		}
	}
}
