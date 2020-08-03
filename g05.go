package main

import (
	"fmt"
)

func main() {
	var guess int
	fmt.Print("Your guess?")
	fmt.Scanf("%v", &guess)
	fmt.Printf("Your guess is %v\n", guess)
}
