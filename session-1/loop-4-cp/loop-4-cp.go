package main

import (
	"fmt"
	//"strings"
)

func createTriangle(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j > n-i {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scanln(&n)
	fmt.Println("Segitiga Siku-siku: ")
	createTriangle(n)
}
