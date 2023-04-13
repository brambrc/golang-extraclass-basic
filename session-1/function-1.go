package main

import (
	"fmt"
)

func luasSegitiga(alas int, tinggi int) int {
	return alas * tinggi / 2
}

func kelilingSegitiga(alas int, tinggi int) int {
	return alas + tinggi + alas
}

func main() {
	var alas int = 10
	var tinggi int = 5

	fmt.Println("Luas Segitiga: ", luasSegitiga(alas, tinggi))
	fmt.Println("Keliling Segitiga: ", kelilingSegitiga(alas, tinggi))

}
