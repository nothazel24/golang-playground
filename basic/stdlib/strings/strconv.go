package basic

import (
	"fmt"
	"strconv"
)

func TestStringConversion() {
	// string to int (return 2 value)
	s := "10"
	if str, err := strconv.Atoi(s); err == nil {
		fmt.Printf("%T\t: %v\n", str, str)
	} else {
		fmt.Println("Konversi error :", err)
	}

	// in to string (hanya return string saja)
	i := 10
	integer := strconv.Itoa(i)
	fmt.Printf("%T\t: %v\n", integer, integer)


	// parse string boolean ke boolean
	resultBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error :", err.Error())
	} else {
		fmt.Println(resultBool)
	}

	// format boolean ke string
	resultStrBool := strconv.FormatBool(false)
	fmt.Printf("%T\t: %v\n", resultStrBool, resultStrBool)

	// NOTE : Untuk penggunaan parse & format bisa dipakai untuk berbagai tipe data
	// jadi, coba ingat ingat dan lihat dokumentasi resminya untuk penggunaan lebih lanjut.
}