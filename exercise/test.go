package exercise

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Menu struct {
	Name  string
	Price int
}

type Order struct {
	Menu     Menu
	Quantity int
}

func (o Order) Subtotal() int {
	return o.Menu.Price * o.Quantity
}

type OrderService struct {
	orders []Order
	menu   map[string]Menu
}

func NewOrderService() *OrderService {
	return &OrderService{
		menu: map[string]Menu{
			"mie ayam":    {Name: "Mie Ayam", Price: 10000},
			"nasi goreng": {Name: "Nasi Goreng", Price: 12000},
			"es teh":      {Name: "Es Teh", Price: 5000},
			"ayam goreng": {Name: "Ayam Goreng", Price: 18000},
		},
	}
}

func (s *OrderService) ShowMenu() {
	fmt.Println("\n=== MENU ===")
	for _, m := range s.menu {
		fmt.Printf("- %s : Rp%d\n", m.Name, m.Price)
	}
}

func (s *OrderService) AddOrder(name string, qty int) error {
	menu, ok := s.menu[strings.ToLower(name)]
	if !ok {
		return fmt.Errorf("menu tidak ditemukan")
	}

	order := Order{
		Menu:     menu,
		Quantity: qty,
	}

	s.orders = append(s.orders, order)
	return nil
}

func (s *OrderService) ShowOrders() {
	fmt.Println("\n=== PESANAN ===")
	total := 0

	for i, o := range s.orders {
		sub := o.Subtotal()
		total += sub

		fmt.Printf("%d. %s x%d = Rp%d\n", i+1, o.Menu.Name, o.Quantity, sub)
	}

	fmt.Printf("Total: Rp%d\n", total)
}

func TestOrderManagement() {
	service := NewOrderService()
	reader := bufio.NewReader(os.Stdin)

	for {
		service.ShowMenu()

		fmt.Print("\nMasukkan nama menu (atau 'exit'): ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if strings.ToLower(name) == "exit" {
			break
		}

		fmt.Print("Jumlah: ")
		var qty int
		fmt.Scanln(&qty)

		err := service.AddOrder(name, qty)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		service.ShowOrders()
	}
}
