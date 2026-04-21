package basic

import "fmt"

func change(o *int, value int) {
	*o = value
}

type Menu struct {
	Nama          string
	Harga, Jumlah int
}

type OrderService struct {
	orders []Menu
}

func (s *OrderService) tambahOrder(nama string, harga, jumlah int) {
	order := Menu{nama, harga, jumlah}
	s.orders = append(s.orders, order)
}

func (s *OrderService) tampilkanOrder() {
	for i, v := range s.orders {
		fmt.Printf("%d. Orderan anda : %s - %d x %d\n", i+1, v.Nama, v.Harga, v.Jumlah)
	}
}

func PointerExample() {
	var numberA int = 4
	var numberB *int = &numberA // disini asterisk (*) digunakan saat mendeklarasikan bahwa variabel ini sebuah pointer

	// (&) Ampersand, digunakan untuk mengambil alamat memori
	// (*) Asterisk, digunakan saat deklarasi pointer, & saat mengakses nilai pointer

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB) // disini asterisk (*) digunakan untuk mengakses nilai
	fmt.Println("numberB (address) :", numberB)

	fmt.Println("==============================")
	fmt.Println("Before\t:", numberA)

	change(&numberA, 12)
	fmt.Println("After\t:", numberA)
	fmt.Println("==============================")

	var slice1 = []string{"Test1", "Test2", "Test3"}
	var testSlice *[]string = &slice1

	*testSlice = append(*testSlice, "Lorem ipsum")
	fmt.Println(slice1)
	fmt.Println(testSlice)

	fmt.Println("==============================")

	service := OrderService{}
	service.tambahOrder("Lorem ipsum", 10000, 3)
	service.tampilkanOrder()
}
