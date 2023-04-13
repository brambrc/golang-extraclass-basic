package main

import (
	"fmt"
)

func add(a int, b int) string {
	result := a + b

	return fmt.Sprintf("Hasil Penjumlahan angka %d dan %d adalah %d: ", a, b, result)
}

func min(a int, b int) string {
	//buat kondisi dimana angka yg lebih besar menjadi pebilang
	result := a - b
	return fmt.Sprintf("Hasil Pengurangan angka %d dan %d adalah %d: ", a, b, result)
}

func div(a int, b int) string {
	//buat kondisi dimana angka yg lebih besar menjadi pebilang
	result := a / b

	return fmt.Sprintf("Hasil Pembagian angka %d dan %d adalah %d: ", a, b, result)
}

func mul(a int, b int) string {
	result := a * b

	return fmt.Sprintf("Hasil Perkalian angka %d dan %d adalah %d: ", a, b, result)
}

func main() {

	var a, b int
	var opt string
	fmt.Print("Masukan Angka Pertama: ")
	fmt.Scanln(&a)
	fmt.Print("Masukan Angka Kedua: ")
	fmt.Scanln(&b)
	fmt.Print("Masukan Operator: ")
	fmt.Scanln(&opt)

	switch opt {
	case "+":
		fmt.Println(add(a, b))
	case "-":
		fmt.Println(min(a, b))
	case "/":
		fmt.Println(div(a, b))
	case "*":
		fmt.Println(mul(a, b))
	default:
		fmt.Println("Operator Tidak Dikenali")
	}

}
