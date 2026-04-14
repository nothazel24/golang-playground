package basic

import "fmt"

func LenCap() {
	var fruits = []string{"Apel", "Mangga", "Anggur", "Jeruk"}
	// LEN & CAP pada slice
	/*

		len() biasanya digunakan untuk menghitung jumlah elemen slice yang ada. Adapun penggunaannya adalah sebagai berikut (ambil data array diatas)

	*/
	fmt.Println(len(fruits))

	/*

		Adapun cap(), biasanya digunakan untuk menghitung capacity / kapasitas maksimum slice. meskipun konsepnya hampur sama dengan penggunaan len(), akan tetapi akan ada perbedaan dalam penggunaannya. berikut penjelasan dan source codenya.

	*/
	laptops := []string{"Thinkpad", "Lenovo", "Asus", "HP", "Alienware"}
	fmt.Printf("Jumlah data yang ditampung: %d\n", len(laptops)) // output 5
	fmt.Printf("Jumlah kapasitas tampungan: %d\n", cap(laptops)) // output 5

	var laptopA = laptops[0:3]
	fmt.Printf("Jumlah data yang ditampung: %d\n", len(laptopA)) // output 3
	fmt.Printf("Jumlah kapasitas tampungan: %d\n", cap(laptopA)) // output 5

	var laptopB = laptops[2:5]
	fmt.Printf("Jumlah data yang ditampung: %d\n", len(laptopB)) // output 3
	fmt.Printf("Jumlah kapasitas tampungan: %d\n", cap(laptopB)) // output 3

	/*

		Mengapa outputnya bisa berubah?
		CONTOH : output len bisa 3 atau output cap bisa 5

		Jawabannya karena output len hanya menghitung data yang ada didalam index, dan cap menghitung kapasitas dari slice itu sendiri. contohnya pada variabel laptopA

		Disana dituliskan untuk me-referensikan data slice dari index ke 0 sampai sebelum index ke-3. yaitu HP. Otomatis akan mengeluarkan output seperti ini:
		[Thinkpad, Lenovo, Asus, ..., ...]

		Ada 2 index kosong disana, nah.., cap akan tetap menghitung kedua index tersebut sebagai kapasitas dari array. sedangkan len hanya menghitung index yang terisi saja.

		berbeda jika index kosong nya ada di sebelum index terisinya. Contoh:
		[..., ..., Asus, HP, Alienware]

		cap() akan menghitung dari index awal didefinisikannya saja (dari index ke 2 dampai akhir), lelu mengeluarkan output 3

	*/
}
