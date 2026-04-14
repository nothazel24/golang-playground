package basic

import (
	"fmt"
	"strconv"
)

func TestDebugging() {
	var input string

	fmt.Print("Type some number : ")
	fmt.Scanln(&input)

	number, err := strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is a number")
	} else {
		fmt.Println(number, "is not a number!")
		fmt.Println(err.Error()) // method Error() akan mengembalikan data dalam bentuk string
	}
}