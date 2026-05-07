package basic

import (
	"fmt"
	"os"
)

func doHelloWorld() string {
	return "Hello world!"
}

func doSomething() {
	fmt.Println("Hari yang cerah bukan?")
}

func getHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	} else {
		return hostname, nil
	}
}

func TestingOS() {
	command := os.Args[1]

	switch command {
	case "hello":
		hello := doHelloWorld()
		fmt.Println(hello)
	case "--do-something":
		doSomething()
	case "--get-hostname":
		result, err := getHostname()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(result)
	default:
		fmt.Println("Perintah tidak dikenali!")
	}

}
