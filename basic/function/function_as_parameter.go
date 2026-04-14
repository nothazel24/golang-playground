package basic

import (
	"fmt"
	"strings"
)

// membuat alias agar penulisan function dalam parameter tidak jadi panjang dan ribet
type Filter func(string) string // hanya menerima function dengan syarat parameter dan return nya string

/*
Function as parameter
Biasanya digunakan untuk membuat function jadi lebih reusable dan fleksibel.

Function utama menerima function lain sebagai argumen (dengan tipe yang sesuai), sehingga behavior / output nya bisa berubah tanpa perlu mengubah kode utama
*/
func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

// function untuk filter ke 1
func spamFilter(name string) string { // ini merupakan function dengan param & return bertipe string
	bannedWords := []string{"Anjing", "Goblok"}

	for _, w := range bannedWords {
		if name == w {
			return strings.Repeat("#", len(w))
		}
	}

	return name
}

// function untuk filter ke 2
func toUpper(name string) string {
	return strings.ToUpper(name)
}

func TestFilter() {
	/*

		Contoh pemanggilannya ada dibawah, jadi kita bisa mengubah output data sesuai dengan function yang kita panggil di argumen dari function utama

	*/

	// dengan menggunakan function filter ke 1
	sayHelloWithFilter("Ryan", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Goblok", filter)

	// dengan menggunakan function filter ke 2
	sayHelloWithFilter("ryan bajindul", toUpper)

}
