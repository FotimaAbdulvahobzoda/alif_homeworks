package main

import "fmt"

func main() {
	var role string

	fmt.Print("Enter your role (admin/manager/user/guest): ")
	fmt.Scan(&role)

	switch role {
	case "admin":
		fmt.Println("Full access available")
	case "manager":
		fmt.Println("User management available")
	case "user":
		fmt.Println("Default access available")
	case "guest":
		fmt.Println("Read-only access available")
	default:
		fmt.Println("Unknown role. Please enter your role")
	}
}
