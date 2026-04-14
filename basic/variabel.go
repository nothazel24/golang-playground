// ini di deklarasikan sesuai dengan nama folder, agar lebih mudah untuk diingat dan digunakan.
package basic

import "fmt"

func TampilkanNama() {
	// variabel
	var nama string = "Ryan"
	umur := 18

	fmt.Println("Nama:", nama)
	fmt.Println("Umur:", umur)

	// if else
	if umur >= 17 {
		fmt.Println("Status: Dewasa")
	} else {
		fmt.Println("Status: Belum dewasa (Remaja)")
	}
}
