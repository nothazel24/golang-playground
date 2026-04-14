package basic

import "fmt"

func CopyExample() {
	/*

		Penggunaan copy() di golang
		Digunakan untuk meng-copy element slice pada src (source / data yang ingin disalin) dan mengirimnya ke dst (destination / tempat tujuan data disimpan)

		contoh:
		copy(dst, src)

	*/

	fruits := []string{"Apel", "Jeruk", "Mangga", "Stroberi"}

	fruits2 := make([]string, len(fruits))
	copy(fruits2, fruits) // copy data slice fruits, dan masukkan ke wadah fruits2

	// insert data slice
	fruits2 = append(fruits2, "Pisang")

	fmt.Println(fruits)  // [Apel, Jeruk, Mangga, Stroberi]
	fmt.Println(fruits2) // [Apel, Jeruk, Mangga, Stroberi, Pisang]
}
