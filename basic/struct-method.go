package basic

import "fmt"

type Customer struct {
	Name, Address string
	age           int
}

func (c *Customer) sayHello(name string) {
	fmt.Println("Halo", name, "nama saya", c.Name)
}

func StructMethodExample() {
	defer func ()  {
		if r := recover(); r != nil {
			fmt.Println("Panic occured : ", r)
		}	
	}()

	budi := Customer{
		Name:    "Budi Doremi",
		Address: "st. Lorem ipsum",
		age:     20,
	}

	budi.sayHello("Ryan bajindul")
}
