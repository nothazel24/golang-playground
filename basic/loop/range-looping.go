package basic

import "fmt"

func RangeLooping() {

	// range looping biasanya digunakan untuk me-looping data gabungan (misalnya string, array, slice, map)
	// NOTES: range akan mengembalikan 2 nilai, jadi buat variabelnya untuk menampung nilai yang dikeluarkan oleh range (Usahakan yang simpel dan mudah diingat)

	fmt.Println("Looping section")

	var xs = "123" // string
	for i, v := range xs {
		fmt.Println("Index:", i, "Values:", v) // index (int), values (str)
	}

	// looping array (hanya value, tanpa index)
	var ys = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range ys {
		fmt.Println("Looping Value:", v)
	}

	fmt.Println("\nSlice section")
	var zs = ys[0:2]
	for _, v := range zs {
		fmt.Println("Slice Value:", v)
	}

	fmt.Println("\nMap section")
	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2}
	for k, v := range kvs {
		fmt.Println("Key=", k, "Value=", v)
	}

	// untuk mengabaikan key dan valuenya (tapi tetap, yang dikeluarkan ada 2 nilai. cuman ada dibelakang layar aja)
	for range kvs {
		fmt.Println("Done")
	}

	// selain itu, bisa juga dengan cukup menentukan nilai numerik perulangan
	fmt.Println("\nNumeric Looping")
	for i := range 5 {
		fmt.Print(i) // 01234
	}
}
