package basic

import "fmt"

func TestClosure() {
	counter := 0

	/*

		Closure
		Merupakan sebuah anonymous function yang ditulis kedalam variabel, dan bisa mengakses data disekitarnya.
		Contoh, closure dibawah bisa mengakses data dari variabel counter diatas

		Ini merupaka fitur yang sangat berguna, karena closure tidak perlu dikirimkan parameter secara terus menerus, namun bisa langsung mengambil data dari luar scope functionnya

		Namun, hati hati saat menggunakan closure, karena dia bisa mengubah data disekitarnya, maka rentan terjadi perubahan data tanpa kamu ketahui. Misal, banyak closure mengakses 1 data secara bergantian / bersamaan. Ini bisa menyebabkan error yang sulit dilacak

	*/
	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
}
