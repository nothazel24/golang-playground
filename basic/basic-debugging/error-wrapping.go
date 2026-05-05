package basic

import (
	"errors"
	"fmt"
)

var (
	errNotFound   = errors.New("Data tidak dapat ditemukan")
	errOutOfRange = errors.New("Melebihi batas wajar")
)

type ValidationError struct {
	Message string
	field   string
}

// error wrapper
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validasi gagal pada field %s. %s", e.field, e.Message)
}

func validateIdentity(name string, age int) error {
	if age < 0 {
		return fmt.Errorf("validateIdentity : %w", &ValidationError{field: "age", Message: "tidak boleh negatif!"})
	}

	if age > 100 {
		return fmt.Errorf("validateIdentity : %w", errOutOfRange)
	}

	if name == "" {
		return fmt.Errorf("validateIdentity : %w", &ValidationError{field: "name", Message: "Nama tidak boleh kosong!"})
	}

	if name != "ryan" {
		return fmt.Errorf("[404] %w", errNotFound)
	}

	return nil
}

func ProcessUser(name string, age int) (string, error) {
	if err := validateIdentity(name, age); err != nil {
		return "", fmt.Errorf("ProcessUser Gagal : %w", err) // %w = wrapper
	}

	return "sbhdcbs", nil
}

func RunErrorWrapping() {
	result, err := ProcessUser("", 19)
	if err != nil {
		fmt.Println(err.Error())

		switch {
		case errors.Is(err, errNotFound):
			fmt.Println(errNotFound)
		case errors.Is(err, errOutOfRange):
			fmt.Println(errOutOfRange)
		default:
			var valErr *ValidationError
			if errors.As(err, &valErr) {
				fmt.Println("Validasi gagal pada field :", valErr.field)
				fmt.Println("Pesan :", valErr.Message)
			} else {
				fmt.Println("Internal server error")
			}
		}
	}

	fmt.Println(result)
}
