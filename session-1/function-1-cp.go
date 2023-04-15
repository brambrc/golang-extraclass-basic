package main

import (
	"fmt"
	"math"
)

func luasLingkaran(jariJari int) float64 {
	return math.Pi * float64(jariJari) * float64(jariJari) //hitung luas lingkaran
}

func kelilingLingkaran(jariJari int) float64 {
	return 2 * math.Pi * float64(jariJari) //hitung keliling lingkaran
}

func main() {
	var jariJari int = 10

	fmt.Println("Luas Lingkaran: ", luasLingkaran(jariJari))
	fmt.Println("Keliling Lingkaran: ", kelilingLingkaran(jariJari))
}
