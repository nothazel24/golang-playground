package basic

import "fmt"

func MapTesting() {
	// Map
	/*

		Kurang lebih konsep penggunaan map sama seperti array dan slice, namun bedanya disini, map menggunakan key dan value untuk mengakses datanya. Key nya sendiri harus ditulis dengan unik, karena digunakan sebagai penanda / identifier untuk pengaksesan value yang ada di map.

		Adapun contohnya adalah sebagai berikut
		map[tipeDataKey]tipeDataValue

	*/
	var month map[string]int // wadah kosong. hanya contoh saja
	month = map[string]int{} // inisialisasi wadah (pake curly bracket{}, agar nilainya bisa ditentukan secara eksplisit)

	// pendefinisan nilai map dari akhir
	month["januari"] = 30
	month["maret"] = 28

	fmt.Printf("Bulan januari ada %d hari\n", month["januari"])
	fmt.Printf("Bulan maret ada %d hari\n", month["maret"])

	// pendefinisian nilai map dari awal
	bulan := map[string]int{"april": 30, "mei": 31}

	delete(bulan, "mei") // hapus map dengan key: "mei" di variabel bulan

	for key, value := range bulan {
		fmt.Printf("Bulan %s ada %d hari\n", key, value) // key: april, value 30
	}

	// deteksi nilai dengan keadaan tertentu
	warna := map[int]string{1: "Biru", 2: "Merah", 3: "Kuning"}
	var value, isExist = warna[1]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Item nya gaada bos..,")
	}

	// Kombinasi map dengan slice
	var chickens = []map[string]string{
		// map[string]string{"name": "Wilson", "gender": "male"},
		// map[string]string{"name": "Wilson2", "gender": "male"},
		// map[string]string{"name": "Wilson3", "gender": "female"},

		// Biar tidak redundan untuk penulisan tipe data map nya
		{"name": "Wilson", "gender": "male"},
		{"name": "Wilson2", "gender": "male"},
		{"name": "Wilson3", "gender": "female"},

		// NOTE : Dalam real case, tiap-tiap elemen bisa saja memiliki key yang berbeda beda. Contoh:
		// {"name": "chicken blue", "gender": "male", "color": "brown"},
		// {"address": "mangga street", "id": "k001"},
		// {"community": "chicken lovers"},
	}

	for _, chicken := range chickens {
		fmt.Printf("Nama ayam: %s, gender: %s\n", chicken["name"], chicken["gender"])
	}
}
