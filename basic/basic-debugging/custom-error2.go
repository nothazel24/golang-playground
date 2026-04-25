package basic

import (
	"fmt"
)

type validationError struct {
	Message string
}

func (v *validationError) Error() string { // sesuai dengan kontrak interface error di golang : Error() string
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string { // ini juga sama
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"data tidak boleh kosong!"}
	}

	if id != "ryan" {
		return &notFoundError{"Data tidak dapat ditemukan"}
	}

	return nil
}

func TestCustomError() {
	err := SaveData("", nil)

	if err != nil {
		// method 1: if, comma ok idiom
		// 	if validationErr, ok := err.(*validationError); ok {
		// 		fmt.Println("Validasi error\t:", validationErr.Message)
		// 	} else if notFoundErr, ok := err.(*notFoundError); ok {
		// 		fmt.Println(notFoundErr.Message)
		// 	} else {
		// 		fmt.Println("Unknown", err.Error())
		// 	}

		// method 2 : switch assertion
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validasi error\t: ", finalError.Error())
		case *notFoundError:
			fmt.Println(finalError.Error())
		default:
			fmt.Println("Unknown error\t: ", finalError.Error())
		}
	}
}
