package basic

import "fmt"

// declare
type User struct {
	Nama string
	Umur int
}

func StructExample() {
	user := User{
		Nama: "Ryan",
		Umur: 18,
	}

	fmt.Println("\nStruct:", user.Nama, user.Umur)
}
