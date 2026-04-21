package basic

import "fmt"

// declare
type User struct {
	Nama string
	Umur int
}

// contoh untuk penggunaan alias
type Person struct {
	nama string
	umur int
}

// nester struct
type student struct {
	person struct { // anonymous struct
		name string
		age  int
	}
	grade   int
	hobbies []string
}

func StructExample() {
	// type Person = orang // alias

	user := User{
		Nama: "Ryan",
		Umur: 18,
	}

	fmt.Println("\nStruct:", user.Nama, user.Umur)

	// w pointer
	var user2 *User = &user
	user2.Nama = "Jamal"

	fmt.Println("=== Dengan pointer ===")
	fmt.Println("Struct:", user.Nama, user.Umur)   // output : jamal 18
	fmt.Println("Struct:", user2.Nama, user2.Umur) // output : jamal 18

	// anonymous struct
	// pendeklarasian struct tanpa menggunakan keyword type akan langsung memasukkan cetakan struct kedalam sebuah variabel
	// ini juga bisa diisi tanpa perlu pendeklarasian secara langsung. contoh :
	// var student struct {
	// 	kelas int
	// }

	// pendeklarasian struct dengan langsung
	var test1 = struct {
		User
		Kelas int
	}{
		// disini bisa digunakan sebagai pendeklarasian nilai langsung menuju struct. Contoh :
		// User: User{"Jamal", 18},
		// Kelas: 4,
	}

	test1.User = User{"wick", 21}
	test1.Kelas = 2

	fmt.Println("name  :", test1.User.Nama)
	fmt.Println("age   :", test1.User.Umur)
	fmt.Println("grade :", test1.Kelas)

	// gabungan antara struct dengan slice
	var testStructLoop = []struct {
		User
		Kelas int
	}{
		{User: User{"Jamalludin", 18}, Kelas: 4},
		{User: User{"Ryan Bajindul", 18}, Kelas: 3},
		{User: User{"Freddy clanker", 999}, Kelas: 90},
	}

	for _, student := range testStructLoop {
		fmt.Println(student.Nama, "Berumur :", student.Umur, "Tahun")
	}
}
