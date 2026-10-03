package main

import (
	"fmt"
	geometry "geo/geo"
)

func main() {
	fmt.Println("----- Rectangle 1 -----")
	rect := geometry.RectangleArea(34.9, 35)
	rect1 := geometry.RectanlePerimeter(54, 67)
	fmt.Printf("Area: %.2f\nPerimeter: %.2f\n", rect, rect1)
	fmt.Println("----- Rectangle 2 -----")
	rect2 := geometry.RectangleArea(55, 67.8)
	rect3 := geometry.RectanlePerimeter(56, 44)
	fmt.Printf("Area: %.2f\nPerimeter: %.2f\n", rect2, rect3)
	fmt.Println("----- Rectangle 3 -----")
	rect4 := geometry.RectangleArea(99, 99)
	rect5 := geometry.RectanlePerimeter(189, 190)
	fmt.Printf("Area: %.2f\nPerimeter: %.2f\n", rect4, rect5)

}
