package main

import "fmt"

type Vehicle interface {
	ShowDetails()
	ShowName() string
}

type Bike struct {
	Name, Color string
	Price       float64
}

func (bike Bike) ShowDetails() {
	fmt.Println("Bike name: ", bike.Name)
	fmt.Println("Bike color: ", bike.Color)
	fmt.Println("Bike price: ", bike.Price)
}

func (bike Bike) ShowName() string {
	return bike.Name
}

func main() {
	var obj Vehicle = Bike{
		Name:  "Honda",
		Color: "Black",
		Price: 100000,
	}
	obj.ShowDetails()
	fmt.Println(obj.ShowName())
}
