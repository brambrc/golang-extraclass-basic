package main

import "fmt"

func main() {

	//program fizzbuzz ini sering sekali dipakai untuk test skill programmer dan test masuk perusahaan
	//tugasnya adalah menampilkan angka dari 1 sampai 15
	//jika angka tersebut habis dibagi 3 maka tampilkan Fizz
	//jika angka tersebut habis dibagi 5 maka tampilkan Buzz
	//jika angka tersebut habis dibagi 3 dan 5 maka tampilkan FizzBuzz
	//jika angka tersebut tidak habis dibagi 3 dan 5 maka tampilkan angka tersebut
	//operasi ini menggunakan operator modulus (%) untuk mecari sisa bagi hasil dan operator logika (&&, ||, !)

	var a int = 15

	for i := 1; i <= a; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
