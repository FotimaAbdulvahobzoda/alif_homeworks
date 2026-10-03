package main

import (
	"fmt"
	catalog "shop/shoplist"
)

func main() {
	products := []catalog.Product{
		{Name: "Shampoo", Price: 8},
		{Name: "Soap", Price: 38},
		{Name: "Toothpaste", Price: 14},
		{Name: "Toothpick", Price: 21},
		{Name: "Shower gel", Price: 26},
		{Name: "Conditioner", Price: 30},
	}
	fmt.Println(products)
	fmt.Println("Total price:", catalog.TotalPrice(products))

}
