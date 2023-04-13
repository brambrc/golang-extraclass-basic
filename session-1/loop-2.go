package main

import "fmt"

func main() {

	var a int
	fmt.Print("Masukan Angka : ")
	fmt.Scanln(&a)

	for x := 1; x <= a; x++ {
		for y := a; y >= x; y-- {
			fmt.Print(" ")
		}

		for z := 1; z <= x; z++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}

}
