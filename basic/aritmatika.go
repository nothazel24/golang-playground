package basic

import "fmt"

func AritmatikaPerbandingan() {
	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual) // output nilai 2 (true)
}
