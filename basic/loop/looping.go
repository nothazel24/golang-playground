package basic

import "fmt"

func NumbersLooping() {
	var j = 0
	var k = 0

	// ini contoh looping dengan menggunakan argumen
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// // looping dengan menggunakan kondisinya saja
	for j < 5 {
		fmt.Println("Angka:", j)
		j++
	}

	// looping tanpa menggunakan argumen
	for {
		fmt.Println("Angka ke -", k) // print dulu

		k++         // tambah value (loop)
		if k == 5 { // baru checking (kondisi)
			break
		}
	}
}
