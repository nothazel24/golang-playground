package basic

import "fmt"

// declare
func tambah(a int, b int) int {
	return a + b
}

// function dengan return multiple
func cekUmur(umur int) (string, bool) {
	if umur >= 17 {
		return "Dewasa", true
	}
	return "Belum dewasa", false
}

func TestFunction() {
	umur := 18

	// memakai function tambah
	hasil := tambah(5,3)
	fmt.Println("\nHasil tambah", hasil)

	// memakai function cek umur
	status, valid := cekUmur(umur)
	fmt.Println("Cek Umur:", status, valid)
}
