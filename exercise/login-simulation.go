package exercise

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	emptyFieldError = errors.New("Field Tidak boleh kosong!")
	notFoundError = errors.New("Data tidak dapat ditemukan")
	ifExistError = errors.New("Data pengguna sudah digunakan !")
)

type Pengguna struct {
	Name, Username, password string
	Age int
}

type Data struct {
	user map[string]Pengguna
}

func NewLoginService() *Data {
	return &Data{
		user: map[string]Pengguna{
			"BUdi90": {Name: "Budi santoso", Username: "BUdi90", password: "jkldkk", Age: 18},
		},
	}
}

func (s *Data) AddPengguna(name, uName, pass string, age int) error {
	if uName == "" || pass == ""  {
		return emptyFieldError
	}

	if _, exists := s.user[uName]; exists {
		return ifExistError
	}

	user := Pengguna{
		Name: name,
		Username: uName,
		password: pass,
		Age: age,
	}

	s.user[uName] = user
	return nil
}

func ExecLoginSimProgram() {
	service := NewLoginService()
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nama Pengguna: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Masukkan username Pengguna: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Masukkan password Pengguna: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	fmt.Print("Masukkan umur Pengguna: ")
	var age int
	fmt.Scanln(&age)

	err := service.AddPengguna(name, username, password, age)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Data berhasil dimasukkan!")
	}
}