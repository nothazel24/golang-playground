package basic

import (
	"flag"
	"fmt"
)

func TestFlagOS() {
	var (
		// flag.T("key", "value", "deksripsi") 
		// value deskripsi biasanya akan ditampilkan ketika kita menulis -h (help)
		// gunanya untuk memberikan penjelasan dari penggunaan key yang ada

		username = flag.String("username", "root", "database username")
		password = flag.String("password", "root", "database password")
		host = flag.String("host", "localhost", "database host")
		port = flag.Int("port", 0, "database port")
	)

	flag.Parse() // digunakan untuk menampung argumen yang dimulai dari index pertama

	// go run . -username="ryan bajindul" -password=rahasia -host=123.234.456.23 -port=5505
	fmt.Println("Username :", *username)
	fmt.Println("Password :", *password)
	fmt.Println("Host :", *host)
	fmt.Println("Port :", *port)
}