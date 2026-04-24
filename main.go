package main

import (
	// import file lain yang ada di folder yang berbeda
	/*
		mengapa diperlukan?
		karena go membaca file per package (1 folder 1 package), jadi jika kamu mendeklarasikannya di folder yang berbeda maka akan dianggap package yang berbeda. maka diperlukan untuk meng-import file yang diperlukan kedalam file utama / file yang membutuhkan file yang ditulis di folder yang berbeda untuk bisa menjalankan kodenya.

		POKOKNYA, BEDA FOLDER = WAJIB IMPORT!
	*/
	// "belajar-go/exercise" // contoh importing file (namaproject/namapackage)
	"belajar-go/basic/interfaces"
)

func main() {
	// exercise.ActivateApp() // contoh pengeksekusian (namapackage.MainFunction())
	// basic.SwitchExample3()
	// basic.TriangleLoop()
	basic.TestEmptyInterface()
	// exercise.TestOrderManagement()
}
