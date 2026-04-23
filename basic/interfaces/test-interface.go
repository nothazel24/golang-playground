package basic

import "fmt"

type Product interface {
	GetCommonInfo() (string, string, string)
	GetSpec() string
	GetPrice() int
}

type Laptops struct {
	Name, Type, SerialNumber string
	price, qty               int
}

type Specification struct {
	Laptops
	Proc    string
	Ram     string
	Storage string
}

func (m *Laptops) GetCommonInfo() (string, string, string) {
	return m.Name, m.Type, m.SerialNumber
}

func (m *Laptops) GetPrice() int {
	return m.price
}

func (m *Specification) GetSpec() string {
	return fmt.Sprintf(
		"%s | %s | %s | %s",
		m.Name, m.Proc, m.Ram, m.Storage,
	)
}

func (m *Specification) GetCommonInfo() (string, string, string) {
	return m.Laptops.GetCommonInfo()
}

func DisplayCommonInfo(value Product) {
	name, typ, sn := value.GetCommonInfo()
	fmt.Printf("Nama: %s. Tipe: %s. Manufaktur: %s\n", name, typ, sn)
}

func DisplaySpecification(value Product) {
	fmt.Println("Spec\t:", value.GetSpec())
	fmt.Println("Harga\t:", value.GetPrice())
}

func Execute() {
	prod := Specification{Laptops: Laptops{
		Name:         "Thinkpad",
		Type:         "T490",
		SerialNumber: "ISF-490",
		price:        6000000,
		qty:          4,
	},
		Proc:    "Intel i5 gen 10",
		Ram:     "16 GB",
		Storage: "260 GB",
	}

	DisplayCommonInfo(&prod)
	DisplaySpecification(&prod)
}
