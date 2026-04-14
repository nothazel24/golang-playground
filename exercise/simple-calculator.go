package exercise

import "fmt"

func tambah(i int, j int) int {
	return i + j
}

func SimpleCalculator() {
	var i, j int

	/*
	
		Gunakan Scan untuk membaca input data. ada 3 tipe scan dalam golang, yaitu Scan, Scanln, dan Scanf

		1. Scan digunakan untuk membaca data yang dipisahkan oleh spasi
		   Contoh input: Ryan 20

		2. Scanln digunakan untuk membaca data sampai enter(newline). Akan berhenti jika user menekan enter
		   Contoh input: Ryan 20
		   
		   Perbedaan dengan Scan biasa:
		   Scanln berhenti di akhir baris
		   Lebih cocok untuk input per baris

		3. Scanf sama konsepnya seperti kedua scan yang sebelumnya, namun dengan menggunakan format khusus seperti bahasa C untuk mengontrol format data apa saja yang bisa dimasukkan ke input. Kalau tidak sesuai dengan format, bisa gagal/error
		   Contoh code: fmt.Scanf("%s %d", &nama, &umur)

		   Note:
		   %s = string
		   %d = int
		   %f = float
		   %v = tampilkan apa adanya (default, tanpa batasan tipe data)
		   %T = cek tipe data
	
	*/ 

	fmt.Println("Masukkan nomor pertama")
	fmt.Scanln(&i)

	fmt.Println("Masukkan nomor kedua")
	fmt.Scanln(&j) // &var = ambil address (alamat memory) dari variabel j

	hasil := tambah(i, j)
	fmt.Println("Hasil penambahan:", hasil)
}
