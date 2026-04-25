package basic

import "fmt"

func random() any {
	return true
}

func random2(flag bool) any {
	if flag {
		return "Testing test"
	} else {
		return 124
	}
}

func TestEmptyInterface() {
	// any / interface{}
	// digunakan untuk mengampung tipe data apapun didalamnya, contoh ada dibawah

	var testMap map[string]any
	var secret any

	// NOTE : Setiap data yang ditulis kedalam interface kosong, apapun tipe datanya
	// biasanya akan ditampilkan sebagai string
	// (jika ingin menampilkannya dengan menggunakan fungsi fmt.Println())

	// jadi jika kita ingin mengolahnya sesuai
	// dengan tipe data aslinya, maka diperlukanlah teknik casting : var.(T)
	// guna mengubah tipe data yang sebelumnya ditampilkan string menjadi tipe data
	// yang kita mau, sehingga sudah siap untuk diolah

	// (!) map dengan key bertipe string dan value bebas
	testMap = map[string]any{
		"name":    "Ryan bajindul",
		"age":     18,
		"hobbies": []string{"Coding", "Designing", "Reading manga", "Hiking"},
	}

	// CONTOH 1
	secret = 2
	var number = secret.(int) * 10 // casting nilai interface
	fmt.Println(secret, "dikali 10 jadi :", number)

	fmt.Println("==============================")

	// CONTOH 2
	fmt.Printf("Nama\t: %s\n", testMap["name"])
	fmt.Printf("Umur\t: %d\n", testMap["age"])
	fmt.Printf("Hobi\t:\n")

	// looping data slice dalam map interface kosong / any
	hobbies := testMap["hobbies"]
	if hobbies, ok := hobbies.([]string); ok { // casting hobbies dalam any menjadi slice string
		for i, hobby := range hobbies {
			fmt.Printf("%d. %s\n", i+1, hobby)
		}
	} else {
		fmt.Errorf("Data tidak tersedia!")
	}

	fmt.Println("==============================")

	// casting interface ke pointer
	type persons struct {
		name string
		age  int
	}

	var person any = &persons{name: "Ryan bajindul", age: 18}

	// ambil data name dari variabel person yang mengacu pada struct persons
	var name = person.(*persons).name
	fmt.Println(name)

	fmt.Println("==============================")

	// kombinasi dengan slice & map pada interface kosong
	var person2 = []map[string]any{
		{"name": "Ryan bajindul", "age": 18},
		{"name": "Abdul dudul", "age": 20},
		{"name": "Asep stroberi", "age": 17},
		{"name": "Yayan kopling", "age": 16},
	}

	for _, each := range person2 {
		fmt.Println(each["name"], "berumur", each["age"])
	}

	fmt.Println("==============================")

	// slice yang isiya bisa apa saja
	var fruits = []any{
		map[string]any{"name": "Strawberry", "total": 10},
		[]string{"Mango", "Apple", "Papaya", "Watermelon"},
		"Orange",
	}

	// ini bisa dibuat lebih kompleks lagi untuk memanipulasi data map, slice, dan array didalamnya
	// untuk mengolah data yang terdapat didalam interface kosong / any
	// contoh nya seperti kode diatas, yang gunanya untuk mengolah data slice dalam interface kosong

	for _, e := range fruits {
		fmt.Println(e)
	}

	fmt.Println("==============================")

	// CASE : Jika terjadi casting dari interface kosong sebanyak 2x
	// dari function random()

	var result any = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// var resultInt int = result.(int) // panic
	// fmt.Println(resultInt)

	// solusi biar tidak panic, pake switch type assertion
	switch value := result.(type) {
	case string:
		fmt.Println("String\t:", value)
	case int:
		fmt.Println("Integer\t:", value)
	default:
		fmt.Println("Unknown")
	}

	// atau, jika hanya ingin mengambil 1 tipe data saja dalam func any
	result2 := random2(true)

	// kita coba untuk mengambil tipe data string saja untuk diolah disini
	str, ok := result2.(string)
	if ok {
		fmt.Println(str)
	} else {
		fmt.Println("Bukan string")
	}
}
