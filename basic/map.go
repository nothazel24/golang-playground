package basic

import "fmt"

func ArrayMapping() {
	userMap := map[string]string{
		"nama": "Ryan bajindul",
		"role": "Admin",
	}

	fmt.Println("\nMap:", userMap["nama"])
}
