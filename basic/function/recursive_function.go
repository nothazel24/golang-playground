package basic

import "fmt"

func recursiveLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	} // 10 * tiap-tiap index (9, 8, 7 - 0)

	return result
}

// ini lebih sederhana dibandingkan menggunakan function dengan looping didalamnya
func recursiveFunction(value int) int {
	if value == 1 {
		return 1 // base case (kondisi berhenti dari recursive function)
	} else {
		/*
			ini akan memanggil functionnya sendiri dengan index yang akan terus dikurangi 1 ketika dipanggil (parameter harus lebih kecil)
		*/
		return value * recursiveFunction(value-1)
	}
}

func TestRecursiveFunction() {
	/*

		Recursive function
		Biasanya digunakan untuk memecah suatu masalah menjadi bagian-bagian terkecil dalam program untuk diselesaikan hingga mencapai base case (kondisi berhenti dari recursive function)

		Recursive function biasa digunakan untuk perhitungan faktorial, Traversal data (DFS), Struktur data bertingkat, Problem solving kompleks, & Algoritma divide dan conquest

		Dalam web dev, biasanya digunakan untuk komentar, sidebar menu karena struktur datanya yang bercabang (tree)
		Komentar A
 		├── Balasan A1
 		│    └── Balasan A1.1
 		└── Balasan A2

		Singkatnya, recursive function digunakan ketika datanya ounya bentuk berulang (self similar), depth / kedalamannya tidak pasti, dan mempunyai struktur bertingkat

		NOTE : Gunakan recursive function dengan bijak, karena recursive function bisa memakan banyak memori dibandingkan dengan loop biasa.

	*/
	fmt.Println(recursiveLoop(10))
}
