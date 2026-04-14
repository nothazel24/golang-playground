package basic

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

/*

	defer function
	function ini memungkinkan kita untuk tetap mengeksekusi program meskipun terjadi error / kesalahan dalam kode function induknya (scope)

	Penulisan defer dilakukan didalam scope function, dan apapun program yang ditambahkan tag defer didalam functionnya, itu akan selalu dieksekusi terakhir / setelah tugas function utamanya selesai.

	ataupun, jika ingin tag defer bisa dieksekusi langsung tanpa perlu menunggu program utama selesai (contohnya untuk mengetest program), kita bisa menggunakan IIFE (Immediately Invoked Function Expression / Function yang langsung dieksekusi)

*/
func runApplication() {
	defer logging()

	// func() { // IIFE
	// 	defer fmt.Println("Testing log...")
	// }()

	fmt.Println("Running application...")
}

func TestDeferFunction() {
	runApplication()
}
