package basic

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Input tidak boleh kosong!")
	}

	return true, nil
}

func TestCustomErrorMessage() {
	var name string
	fmt.Print("Masukkan nama anda : ")
	fmt.Scanln(&name)

	// tanpa if with inilization statement
	// valid, err := validate(name)
	// if (valid) {
	// 	fmt.Println("Halo, ", name)
	// } else {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(err.Error()) // variabel valid dan err masih bisa dipakai dibawah sini

	if valid, err := validate(name); valid { // if with inilization statement (go idiomatic)
		fmt.Println("Halo, ", name)
	} else {
		fmt.Println(err.Error())
	}

	// if with inilization statement
	// memungkinkan kita untuk membuat variabel yang hanya hidup dalam scope if else saja
	// scope selain itu, tidak dapat mengakses variabelnya
	// ini ide

	// fmt.Println(err.Error()) // compiler akan kebingungan, karena tidak ada variabel err untuk diakses disini
}
