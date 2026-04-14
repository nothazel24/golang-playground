package basic

import "fmt"

func ArrayExample() {

	// array strict / defined length
	var angka [3]int = [3]int{1, 2, 3}

	// penggabungan
	fmt.Println("\nArray", angka)

	// slice (dynamic array :=)
	buah := []string{"apel", "jeruk"}
	buah = append(buah, "mangga")

	// penggabungan
	fmt.Println("Slice:", buah) // untuk pemanggilan 1 array, cukup gunakan [index] di sebelah variabel

	// loop slice
	fmt.Println("\nLoop buah")
	for index, value := range buah {
		// range, secara default mengembalikan nilai: index, dan value. jadi jika kamu cuma mengambil 1 variabel maka go akan mengembalikan index saja, bukan value.

		// jika ingin mengambil value nya saja, maka gunakan _ pada index, yang fungsinya untuk mengabaikan index pada range. contoh

		// for _, value := range buah

		// call
		// fmt.Println(value)

		fmt.Println(index, value)
	}
}
