package main

import "fmt"

func main() {
	age()
	temperature()
	myCity()
	fmt.Println("====== Fourth task ======")
	energy := 100
	fmt.Println("Energy left(First time): ", spendEnergy(&energy, 50))
	fmt.Println("Energy left(Second time): ", spendEnergy(&energy, 20))
	fmt.Println("Energy left(Last time): ", spendEnergy(&energy, 110))
	fmt.Println("====== Fifth task ======")
	num := 9
	result := double(num)
	fmt.Println("without pointer")
	fmt.Printf("Before: %v\nAfter: %v\n", num, result)
	fmt.Println("-----------------")
	fmt.Println("with pointer")
	doublePtr(&num)
	fmt.Println("The result: ", num)
	fmt.Println("====== Sixth task ======")
	salar := 3500.0
	applyBonus(&salar, 650.0)
	fmt.Println("Waige with bonus: ", salar)
	fmt.Println("====== Seventh task ======")
	name1, name2 := "Vivnie", "Katherine"
	fmt.Println("Before: ", name1, name2)
	swapStrings(&name1, &name2)
	fmt.Println("After: ", name1, name2)
	fmt.Println("====== Eighth task ======")
	users := map[string]string{
		"Alisher": "Dushanbe",
		"Alex":    "New York",
		"Masha":   "Moscow",
	}
	city := findCity(users, "Masha")

	if city != nil {
		fmt.Println("City:", *city)
	} else {
		fmt.Println("User not found")
	}
	city = findCity(users, "Daniel")
	if city != nil {
		fmt.Println("City:", *city)
	} else {
		fmt.Println("User not found")
	}
	fmt.Println("====== Ninth task ======")
	fmt.Println("The result: ", safeDivide(2, 5))
	fmt.Println("The result: ", safeDivide(10, 2))
	fmt.Println("The result: ", safeDivide(0, 5))
	fmt.Println("The result: ", safeDivide(9, 0))
	fmt.Println("====== Tenth task ======")
	pInt := new(int)
	*pInt = 100
	fmt.Println(*pInt)
	fmt.Println("====== Eleventh task ======")
	array := [5]int{5, 10, 15, 20, 25}
	fmt.Println("Before: ", array)
	doubleArray(&array)
	fmt.Println("After: ", array)
	fmt.Println("====== Twelth task ======")
	a := 1
	b := 2
	c := 3
	candidates := []string{"Aisha", "Kattie", "Matthew", "Anastasiya", "Aisha", "Anastasiya", "Tom", "Aisha", "Kattie", "Matthew", "Anastasiya", "Aisha", "Anastasiya", "Tom", "Aisha", "Kattie", "Matthew", "Anastasiya", "Aisha", "Anastasiya", "Tom"}
	for _, can := range candidates {
		tallyVote(can, &a, &b, &c)
	}
	fmt.Println("Aisha: ", a, "Anastasiya: ", b, "Matthew: ", c)
	fmt.Println("====== Thirteenth task ======")
	temperature := 36.6
	min := 30.0
	max := 45.0
	limited := adjustTemperature(&temperature, 7, min, max)
	fmt.Println("Current temperature: ", temperature, "Limited?: ", limited)
	limited1 := adjustTemperature(&temperature, 10, min, max)
	fmt.Println("Current temperature: ", temperature, "Limited?: ", limited1)
	limited2 := adjustTemperature(&temperature, -8, min, max)
	fmt.Println("Current temperature: ", temperature, "Limited?: ", limited2)

}

func age() {
	age := 20
	fmt.Println("====== First task ======")
	fmt.Printf("The variable: %v\nThe adress of the variable: %v\nThe type of the variable: %T\n", age, &age, age)
}

func temperature() {
	temp := 36
	t := &temp
	*t -= 5
	fmt.Println("====== Second task ======")
	fmt.Printf("Before: %v\nAfter: %v\n", temp, *t)
}

func myCity() {
	city := "Dushanbe"
	pointer := &city
	*pointer = "Istanbul"
	fmt.Println("====== Third task ======")
	fmt.Println("The city: ", *pointer)
}

func spendEnergy(energy *int, cost int) int {
	if *energy <= cost {
		*energy = 0
	} else {
		*energy -= cost
	}
	return *energy
}

func double(n int) int {
	n *= 2
	return n
}

func doublePtr(n *int) {
	*n *= 2
}

func applyBonus(salary *float64, bonus float64) {
	*salary += bonus
}

func swapStrings(a, b *string) {
	*a, *b = *b, *a
}

func findCity(users map[string]string, login string) *string {
	city, ok := users[login]

	if !ok {
		return nil
	}

	return &city
}

func safeDivide(a, b int) *int {
	if a/b == 0 || a == 0/b {
		return nil
	}
	divide := a / b
	return &divide
}

func doubleArray(arr *[5]int) {
	for i := 0; i < len(*arr); i++ {
		arr[i] *= 2
	}
}

func tallyVote(candidate string, countA, countB, countC *int) {
	if candidate == "Aisha" {
		*countA += 1
	} else if candidate == "Anastasiya" {
		*countB += 1
	} else if candidate == "Matthew" {
		*countC += 1
	}
}

func adjustTemperature(current *float64, delta, min, max float64) bool {
	*current += delta

	if *current > max {
		*current = max
		return true
	}

	if *current < min {
		*current = min
		return true
	}

	return false
}
