package basic

import (
	"fmt"
	"strings"
)

func TestFunction() {
	// tanpa variadic function
	var names = []string{"Ryan", "Ferdiansyah"}
	ExampleFunction("Halo", names) // penulisan function & argumen / data yang akan dimasukkan dalam parameter

	// dengan variadic function
	greetings("Halo", "Ryan ", "B. ", "Ferdiansyah")

	/*

		konversi data slice untuk dimasukkan kedalam variadic function (jika sudah terlanjur memiliki variabel untuk dimasukkan kedalam variadic function)

		tinggal masukkan ... kedalam nama argumen. contoh (names...)

	*/
	greetings("Hai", names...)

	fmt.Println(greetings2("Ryan ", "Ferdiansyah"))

}

// function tanpa return nilai / data
func ExampleFunction(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

// variadic function
/*

	function yang menampung banyak data dalam 1 parameter. func namafunction (arr ...string)

	fungsinya untuk mempermudah penulisan function saja. agar tidak perlu mendeklarasikan data slice lagi didalam variabel seperti contoh diatas (dalam TestFunction), atau menampung data dalam tipe slice didalam parameter []dataType

	NOTE: variadic function harus diletakkan di akhir paramater!

*/
func greetings(message string, arr ...string) {
	var nameString = strings.Join(arr, "")
	fmt.Println(message, nameString)
}

// function dengan pengembalian data (return)
/*

	Ini biasanya digunakan untuk memproses suatu data dan mengembalikan data nya untuk dipeoses kembali di main function / function yang lain.

*/
func greetings2(firstName, lastName string) string {
	fullName := firstName + lastName
	return fullName
}

// named return variable
/*

	Digunakan untuk mempermudah penulisan kode, agar pengembalian nilainya tidak lagi dituiskan setiap kali me-return data dalam function, tapi didefinisikan dalam blok pengembalian data setelah parameter namaFunction(param) (dataReturn T)

	NOTE : ini juga bisa mengembalikan nilai lebih dari satu (multiple return variable)

*/
func NamedValue(firstName, lastName string) (fullName string) {
	fullName = firstName + lastName
	return // dari sini, kita tidak perlu menuliskan nilai mana yang harus dikembalikan
}
