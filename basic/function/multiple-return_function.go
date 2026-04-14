package basic

import (
	"fmt"
	"math"
)

func MultipleFunction() {
	var diameter float64 = 15

	/*

		function as a value

		di golang function dianggap sebagai tipe data juga, jadi function bisa menjadi sebuah value didalam variabel. Ini memungkinkan kita untuk memasukkan function sebagai argumen kedalam parameter function yang lain.

	*/
	var area, circumference = calculate(diameter) // dalam kasus ini, function digunakan untuk menampilkan data yang sudah dikembalikan oleh function calculate()

	fmt.Printf("Luas lingkaran\t\t: %2.f\n", area)            // proses 1 : cetak luas lingkaran
	fmt.Printf("keliling lingkaran\t: %2.f\n", circumference) // proses 2 : cetak keliling lingkaran
}

func calculate(d float64) (area, circumference float64) { // mengembalikan 2 nilai
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	return
}
