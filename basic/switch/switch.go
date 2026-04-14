package basic

import "fmt"

func SwitchExample() {
	var score int = 3

	switch score {
	case 8:
		fmt.Println("Perfect!")
	case 7, 6, 5, 4: // 1 case bisa diisi oleh berbagai kondisi sekaligus
		fmt.Println("Awesome!")
	default: // default juga sama
		{
			fmt.Println("You're not bad")
			fmt.Println("You can be better!")
			fmt.Print("Yahahahaha.., hayyyuk...")
		}
	}
}