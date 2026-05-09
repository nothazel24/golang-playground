package basic

import (
	"fmt"
	"strings"
	"unicode/utf8"
)


// pelajari lebih dalam jika sedang luang btw (lumayan penting nih)
// rune, utf8
func PotongString(t string, limit int) string {
	// jika lebih panjang dari limit
	if utf8.RuneCountInString(t) <= limit {
		return t
	}

	// konversi ke rune
	runes := []rune(t)

	// cut & return ...
	return string(runes[:limit]) + "..."
} 

func TestString() {
	fmt.Println(strings.Contains("Ryan ferdiansyah", "yan")) // cari apakah data mengandung kata yang diinginkan?
	fmt.Printf("%q\n", strings.Split("Ryan,Ferdiansyah", ",")) // memisahkan data dan menyimpannya dalam slice
	fmt.Println(strings.ToLower("RYAN FERDIANSYAH"))
	fmt.Println(strings.ToUpper("ryan ferdiansyah"))
	fmt.Println(strings.Trim("     Ryan ferdiansyah     ", " "))
	fmt.Println(strings.ReplaceAll("Ryan ferdiansyah, Ryan bajindul", "Ryan", "Budi"))

	data := "apel,mangga,anggur,rambutan"
	parts := strings.Split(data, ",")

	for i, v := range parts {
		fmt.Println(i+1, ".", strings.TrimSpace(v))
		// strings.TrimSpace(value)
		// sama seperti trim, namun ini tidak perlu menuliskan mana yang harus dipotong lagi
	}

	// experiment
	data2 := "Lorem ipsum dolor sit amet Constectieur lorem amet sit ipsum"
	fmt.Println(PotongString(data2, 31))

	data3 := "Ryan bajindul"
	test := strings.Contains(data3, "bajindul")
	if test {
		fmt.Println("Rill cuy")
	} else {
		fmt.Println("Fek cuy..")
	}

}
