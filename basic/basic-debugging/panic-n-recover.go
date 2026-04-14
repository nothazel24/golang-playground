package basic

import (
	"errors"
	"fmt"
	"strings"
)

func validateName(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Input tidak boleh kosong!")
	}

	return true, nil
}

func catch() {
	// recover()
	// digunakan untuk menghandle panic error
	// pada saat panic error terjadi, recover meng-take over goroutine yang sedang panic supaya program tidak crash total

	// NOTE : recover() hanya bekerja jika dipanggil dalam defer di goroutine yang sama,
	// dan dalam stack panic yang sama. Cntohnya ada dalam function TestPanic() dibawah
	if r := recover(); r != nil { // recover
		fmt.Println("[System] Terjadi kesalahan : ", r)
	} else {
		fmt.Println("[System] Aplikasi berjalan dengan sempurna")
	}
}

func TestPanic() { // function ini memiliki panic, dan memanggil recovery dengan menggunakan defer
	defer catch() // catch panic dengan recover()

	var name string
	fmt.Print("Masukkan nama anda : ")
	fmt.Scanln(&name)

	if valid, err := validateName(name); valid {
		fmt.Println("Halo,", name)
	} else {
		defer fmt.Println("End")
		panic(err.Error()) // akan menghentikan seluruh program dalam scope goroutine / function
	}
}

func IIFEPanic() {
	data := []string{"mobil", "motor", "sepeda"}

	// PANIC ISOLATION (sebuah pattern / pola penanganan masalah)
	for _, each := range data {
		func() { // mini scope per iterasi

			// mini scope per iterasi berguna untuk membuat function baru
			// ketika function yang lama dihentikan prosesnya karena terjadi panic
			// maka, diiterasi selanjutnya akan membuat scope function baru dan menjalankan program yang ada didalamnya
			// dan, jika terjadi panic lagi, maka akan mengulang siklus yang sebelumnya

			defer func() { // recover (membantu menjalankan program kembali ketika terjadi panic)
				if r := recover(); r != nil {
					fmt.Println("Terjadi panic saat perulangan berlangsung", each, "| pesan : ", r)
				} else {
					fmt.Println("Aplikasi berjalan dengan sempurna")
				}
			}()

			// defer fmt.Println("Terjadi panic saat perulangan berlangsung", each)

			// panic ini hanya akan menghentikan function IIFE saja
			// dan langsung lanjut ke iterasi selanjutnya serta mengulang pola yang sama sampai data dari perulangan habis
			panic("Terjadi error yang tak terduga")
		}()
	}
}
