package basic

import (
	"bytes"
	"fmt"
	"os"
)

func fmtExample() {
	fmt.Println("Hello world!")                // langsung ditampilkan diterminal
	str := fmt.Sprintln("Hello", "World", "!") // disimpan dalam bentuk string terlebih dahulu
	fmt.Println(str)

	// Fprint
	// digunakan untuk write ke file / buffer / writer apapun
	age, hobby := 18, "coding"
	file, _ := os.Create("output.txt")
	fmt.Fprintln(file, "Hello there!")
	fmt.Fprintln(file, "My name is ryan ferdiansyah")
	fmt.Fprintf(file, "My age is %d, and my hobby is %s", age, hobby)

	// buffer
	// digunakan untuk mengatur kapan data bisa dikirim lewat network / http
	// data dalam buffer biasanya disimpan terlebih dahulu dalam RAM, dan bebas mau diunakan kapan saja
	var buff bytes.Buffer
	fmt.Fprintf(&buff, "Total %d", 100)
	fmt.Println(buff.String())

	// Stdout
	fmt.Fprintln(os.Stdout, "ini stdout")

	test := fmt.Sprintf("%[2]d, %[1]d\n", 11, 30) // secara eksplisit membalikkan urutan print %[idx]T
	fmt.Println(test)

	// testing
	const nama, umur = "ryan", 18
	n, err := fmt.Fprintf(os.Stdout, "%s is %d years old.\n", nama, umur)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf\t: %v\n", err)
	}

	fmt.Printf("%d bytes written\n", n)
}
