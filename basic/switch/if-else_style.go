package basic

import "fmt"

func SwitchExample2() {
	var score = 6

	/*
		Secara default, switch akan mengehentika proses dan mengeluarkan output jika kondisinya terpenuhi
	*/
	switch {
	case score == 8:
		fmt.Println("Perfect!")
	case (score < 8) && (score > 3):
		fmt.Println("Awesome!")
		/*
			fallthrough:

			akan meneruskan untuk mengeksekusi switch selanjutnya. tidak peduli apakah kondisinya sudah terpenuhi atau tidak
		*/
		fallthrough
	case score < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("You're not bad")
			fmt.Println("You can be better!")
			fmt.Print("Yahahahaha.., hayyyuk...")
		}
	}
}

// switch bersarang (if else)
func SwitchExample3() {
	var point = 0

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("Perfect!")
		default:
			fmt.Println("Nice")
		}
	} else {
		if point == 5 {
			fmt.Println("Not bad..")
		} else if point < 5 && point >= 3 {
			fmt.Println("Keep trying")
		} else {
			fmt.Println("You can do it!")
			
			if point == 0 {
				fmt.Println("Hell nah...")
			}
		}
	}
}
