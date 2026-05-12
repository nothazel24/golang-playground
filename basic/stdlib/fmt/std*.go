package basic

import (
	"bufio"
	"fmt"
	"os"
)

func stdErrOutExample() {

	/*
	
		Stderr & out
		Biasanya dipakai untuk logging dalam pengembangan aplikasi dan digabung dengan Fprint*
		
		Stderr & out biasanya dipakai secara tidak langsung oleh kita, dengan menggunakan stdlib
		yang menerapkan konsep yang serupa, namun penulisannya lebih simple dan mudah
	
	*/

	// stdout / standard out
	// merupakan program untuk menampilkan output kedalam terminal
	// fungsinya sama seperti fmt.Print*(), namun ini lebih eksplisit lagi
	os.Stdout.WriteString("Hello world!\n")
	fmt.Fprintln(os.Stdout, "Hey there!")

	// stderr / standard error
	fmt.Fprintln(os.Stderr, "Oopps.., Terjadi kesalahan. Tidak ada file yang terdeteksi")

	// stdin / standard input
	fmt.Print("input\t: ")
	scanner := bufio.NewScanner(os.Stdin) // bufio.NewScanner digunakan untuk membaca seluruh input
	if scanner.Scan() {
		// display input
		fmt.Println("Anda memasukkan\t:", scanner.Text())
	}
}
