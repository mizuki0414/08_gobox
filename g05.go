package main

import (
	"fmt"
)

func main() {
	var guess int
	fmt.Print("Your guess?")
	// Scanfが標準入力になる話
	fmt.Scanf("%v", &guess)
	fmt.Printf("Your guess is %v\n", guess)
}
