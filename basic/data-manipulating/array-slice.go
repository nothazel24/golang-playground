package basic

import "fmt"

func ArraySlice() {
	// CAPACITY (TANPA PRE-ALLOCATION)
	// var numbers []int

	// DENGAN PRE-ALLOCATION
	// numbers := make([]int, 0, 1000)

	/*

		Kenapa pre-alokasi diperlukan?
		Karena, jika tidak array akan melakukan realokasi / reallocation memori (biasanya sekitar 2x lipat dari nilai awal) dan akan mengcopy data lama ke blok memory baru yang 2x lebih besar. Ini bisa menyebabkan penurunan performa pada aplikasi

		Karena itu make([]T, length, capacity) diperlukan agar kapasitas data sudah ditetapkan dari awal, dan menghindari banyak realokasi / reallocation

	*/

	testArray := [4]string{
		"Apple",
		"Orange",
		"Grape",
		"Pear",
	}

	fmt.Printf("Jumlah elemen: %d\n", len(testArray))
	fmt.Printf("Isi semua elemen: %s\n\n", testArray)

	fmt.Println("==================================")

	var insertArray [4]string
	insertArray[0] = "Testing 1"
	insertArray[1] = "Testing 2"
	insertArray[2] = "Testing 3"
	insertArray[3] = "Testing 4"
	// insertArray[4] = "Testing 5" // ini bakalan error, karena array telah di set untuk menampung 4 nilai saja sebelumnya.

	var numbers = []int{1, 2, 3, 4, 5, 5}

	fmt.Println(insertArray)
	fmt.Printf("Array: %d\n", numbers)
	fmt.Printf("Jumlah elemen: %d\n", len(numbers))

	// array multidimensi
	/*
		[2] menandakan kapasitas baris array
		[3] menandakan kapasitas kolom array

			Contoh:
			[1 2 3]
			[4 5 6]

		2 baris, dan tiap 1 baris memiliki 3 kolom / nilai
	*/
	var multiArray = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	// NOTE: untuk kapasitas array multi dimensi juga bisa dibiarkan tanpa kapasitas awal, jadi data bisa dinamis

	fmt.Println(multiArray)

	// SLICE
	// Pengaksesan Elemen Slice dengan 3 index
	/*

		3 index adalah teknik slicing untuk pengaksesan elemen yang sekaligus menentukan kapasitasnya. Cara penggunaannya yaitu dengan menyisipkan angka kapasitas di belakang, seperti fruits[0:1:1]. Angka kapasitas yang diisikan tidak boleh melebihi kapasitas slice yang akan di slicing.

	*/

	var fruits = []string{"apple", "grape", "banana"}
	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2]

	fmt.Println(fruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits)) // len: 3
	fmt.Println(cap(fruits)) // cap: 3

	fmt.Println(aFruits)      // ["apple", "grape"]
	fmt.Println(len(aFruits)) // len: 2
	fmt.Println(cap(aFruits)) // cap: 3

	fmt.Println(bFruits)      // ["apple", "grape"]
	fmt.Println(len(bFruits)) // len: 2
	fmt.Println(cap(bFruits)) // cap: 2
}
