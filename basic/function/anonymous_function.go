package basic

import (
	"fmt"
	"strings"
)

var bannedWords = []string{"Anjing", "Goblok"}

type Blacklist func(string) bool

func censor(name string) string {
	for _, w := range bannedWords {
		if strings.ToLower(name) == strings.ToLower(w) {
			return strings.Repeat("#", len(w))
		}
	}

	return name
}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You're blocked", censor(name)+"!") // output yang diharapkan : You're blocked, ######!
	} else {
		fmt.Println("Welcome", name)
	}
}

func RunningAnonymousFunction() {
	// cara pertama
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("Ryan", blacklist)

	// Cara kedua
	registerUser("anjing", func(name string) bool {
		for _, w := range bannedWords {
			if strings.ToLower(name) == strings.ToLower(w) {
				return true
			}
		}

		return false
	})

	// IIFE
	// adalah teknik menjalankan Anonymous function secara langsung saat dideklarasikan
	// (tidak perlu memanggilnya ketika ingin menjalankan functionnya)
	// biasanya digunakan untuk meng-enkapsulasi variabel, inisialisasi, atau membatasi scope variabel
	// supaya tidak mencemari global scope. dan juga, IIFE biasa digunakan dengan defer
	func(message string) { // parameter
		fmt.Println(message) // process / logic
	}("Hello world!") // value
}
