package main

import (
	"fmt"
)

func main() {

	var a int
	fmt.Print("Mari Belajar Berhitung : ")
	fmt.Scanln(&a)

	for i := 1; i <= a; i++ {
		fmt.Println("Ini adalah angka ke ", i)
	}
}
