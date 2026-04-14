package basic

import "fmt"

func TriangleLoop() {
	for i := 4; i >= 0; i-- { // looping spasi
		for j := i; j < 5; j++ { // looping
			fmt.Print(j, " ")
		}

		fmt.Println()
	}
}

func Pyramid() {
	n := 5

	for b := 1; b <= 2; b++ {
		for i := 1; i <= n; i++ {
			// spasi
			for j := 1; j <= n-i; j++ {
				fmt.Print(" ")
			}

			for k := 1; k <= i; k++ {
				fmt.Print("* ")
			}

			fmt.Println()
		}
	}
}
