package basic

import "fmt"

// interface
/*

	Interface adalah tipe data di Go yang digunakan untuk mendefinisikan "kontrak perilaku" (behavior).

	Artinya:
	- Interface hanya berisi deklarasi method (tanpa implementasi).
	- Suatu struct akan otomatis dianggap sebagai bagian dari interface
  	  jika ia memiliki method yang sesuai dengan kontrak tersebut.
  	  (tanpa perlu keyword seperti "implements" seperti di OOP lain)
	- Interface digunakan agar suatu function tidak bergantung pada tipe data tertentu (generic)

	Catatan penting:
	- Interface bukan wadah untuk menyimpan method,
  	  melainkan hanya mendeskripsikan method yang harus dimiliki.
	- Method tetap dimiliki dan diimplementasikan oleh struct.

*/
type HasName interface {
	GetName() string // nama method, return (kalau ada, bisa juga jadi multiple return)
}

// struct & method
type Person struct {
	Name string
}

/*
Method ini dimiliki oleh struct Person.

Karena method ini sesuai dengan kontrak interface HasName,
maka secara otomatis Person dianggap mengimplementasikan HasName.
*/
func (n *Person) GetName() string {
	return n.Name
}

/*
Function ini menerima parameter bertipe interface (HasName).

Artinya:
- Function tidak peduli tipe asli (Person, dll)
- Selama memiliki method GetName(), maka bisa digunakan

Di dalam function, kita hanya bisa mengakses method
yang didefinisikan di interface.
*/
func SayHello(value HasName) {
	fmt.Println("Hello,", value.GetName()) // disini wajib untuk memanggil method yang ada di interface
}

func InterfaceExample() {
	person := Person{Name: "Ryan bajindul"}

	// Mengirim pointer karena method GetName menggunakan receiver *Person
	SayHello(&person)
}
