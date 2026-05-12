package basic

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Writer
/*

	Merupakan sebuah interface dalam stdlib io, yang gunanya untuk menulis / mencetak teks kemana saja
	Bisa ke :
	terminal (os.Stdout)
	ke file (os.File)
	ke http response (http.ResponseWriter)
	dan yang terakhir ke memori / RAM (bytes.buffer)

*/

func sendMessage(w io.Writer) {
	w.Write([]byte("Hello there!, my name is ryan\n"))
}

func writerBufferExample() {
	sendMessage(os.Stdout) // terminal

	file, _ := os.Create("test-output2.txt")
	sendMessage(file) // file

	// sendMessage(w) // http response (contoh kecil)

	// Buffer
	// digunakan untuk mengatur kapan data bisa dikirim lewat network / http
	// data dalam buffer biasanya disimpan terlebih dahulu dalam RAM, dan bebas mau digunakan kapan saja
	var testBuffer bytes.Buffer // ditulis ke memori / RAM
	fmt.Fprintln(&testBuffer, "Excuse me, what??")
	fmt.Println(testBuffer.String())
}
