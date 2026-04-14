package basic

import "fmt"

func ArrayLoop() {
	var buah = []string{"Mangga", "Anggur", "Melon", "Pisang"}

	// loop biasa
	for i := 0; i < len(buah); i++ {
		fmt.Printf("Elemen %d : %s\n", i, buah[i])
	}

	// range loop (modernize loop)
	for j, fruit := range buah {
		fmt.Printf("Elemen %d : %s\n", j, fruit)
	}
}
