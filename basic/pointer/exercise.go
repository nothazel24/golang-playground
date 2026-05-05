package basic

import (
	"errors"
	"fmt"
)

type Orang struct {
	Nama, Alamat, Status string
	Umur                 int
}

type OrangService struct {
	Data []Orang
}

func (s *Orang) GetNama() string {
	return s.Nama
}

func (s *Orang) GetIdentitas(nama string) (*Orang, error) {
	if nama != s.Nama {
		return nil, errors.New("Tidak ada nama yang sesuai")
	}

	orang := &Orang{
		Nama:   nama,
		Alamat: s.Alamat,
		Status: s.Status,
		Umur:   s.Umur,
	}

	return orang, nil
}

func (s *OrangService) AddOrang(nama, alamat, status string, umur int) error {
	if nama == "" {
		return errors.New("Nama tidak boleh kosong")
	}

	s.Data = append(s.Data, Orang{
		Nama:   nama,
		Alamat: alamat,
		Status: status,
		Umur:   umur,
	})
	
	return nil
}

func (s *OrangService) GetAll() []Orang {
	return s.Data
}

func PointerExercise() {
	orangService := &Orang{
		Nama:   "Ryan bajindul",
		Alamat: "Kp. Gugunungan",
		Status: "Pelajar",
		Umur:   19,
	}

	service := &OrangService{}

	nama := orangService.GetNama()
	fmt.Println(nama)
	
	identitas, err := orangService.GetIdentitas("Ryan bajindul")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(identitas)

	service.AddOrang("Ryan", "St. Lorem", "Pelajar", 20)

	data := service.GetAll()

	for i, v := range data {
		fmt.Printf("%d. %s, %s, %s - %d\n", i+1, v.Nama, v.Alamat, v.Status, v.Umur)
	}
}
