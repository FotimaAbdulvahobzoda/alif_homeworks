package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	printWelcome()
	printDivider()
	showMenu()
	announceGameOver()
	fmt.Printf("Total: %v", discountPrice(300, 50))
	fmt.Printf("Is he/she an adult?: %v", isAdult(14))
	fmt.Printf("The max of three: %v", maxOfThree(4, 7, 9))
	fmt.Printf("The body mass index: %v",bmi(161.7, 52.5))
	fmt.Println("Minute: %v",secondsToMinutes(360))
	fmt.Printf("Full name: %v", fullName("Fotima", "Abdulvahobzoda"))
	division, modulo := divmod(10, 5)
	fmt.Printf("Division: %v\nModulo: %v\n",division, modulo)
	rectangle, parameter := RectangleParams(34, 20)
	fmt.Printf("Area: %v\nParameter: %v\n", rectangle, parameter)
	name, name1 := parseFullName("Fotima Abdulvahobzoda")
	fmt.Printf("Name: %s, First name: %s\n", name, name1)
	password, boolean := checkPassword("1234567890")
	fmt.Printf("checking %v ....", password, boolean)
	xtotal, ytotal := midpoint(4.6, 7.2, 5.6, 3.9)
	fmt.Printf("The midpoint: (%v: %v)",xtotal, ytotal)
	printSlice({3, 6, 10, 15, 26, 34, 50})
	sliceLength({3, 6, 10, 15, 26, 34, 50})
	makefruits()
	addToSlice({3, 6, 10, 15, 26, 34, 50, 58, 60, 67}, 90)
	printPrices({
		"milk":           7,
		"cheese":         6,
		"yogurt":         4,
		"cottage cheese": 10,
	})
	fmt.Println(makeScores())
	fmt.Println("Applied discounts to the silce: %v", applyDiscount({100, 300, 560, 900, 1200}))
	fmt.Println(wordFrequency())
	fmt.Printf("Even numbers: %v", filterEven({0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}))
	expename, expenprice := mostExpensive({
		"milk":           7,
		"cheese":         6,
		"yogurt":         4,
		"cottage cheese": 10,
	})
	fmt.Printf("Name: %v\n, Price: %v\n", expename, expenprice)
	fmt.Println(averageGrade())
	letters1, digits1 := countLettersAndDigits("Hello12345")
	fmt.Printf("Letters; %v\n, Digits: %v\n", letters1, digits1)


}

func printWelcome() {
	fmt.Println("Welcome to the application!")
}

func printDivider() {
	fmt.Println("=========================")
}

func showMenu() {
	menu := []string{"Expresso", "Latte", "Tea", "Mohito"}
	for il := 1; il <= len(menu); il++ {
		for _, drink := range menu {
			fmt.Sprintf("%v. %s", il, drink)
		}
	}
}

func announceGameOver() {
	fmt.Println("GAME OVER")
	fmt.Println("*************************")
}

func discountPrice(price, percent float64) float64 {
	percent /= 100
	finalPrice := price - (percent * price)
	return finalPrice
}

func isAdult(age int) bool {
	if age >= 18 {
		return true
	} 
	return false
	
}

func maxOfThree(a, b, c int) int {
	max := 0
	combined := []int{a, b, c}
	for _, combined := range combined {
		if combined > max {
			max = combined
		}
	}
	return max
}

func bmi(heightM, weightKg float64) float64 {
	index := weightKg / (heightM * heightM)
	return index
}

func secondsToMinutes(seconds int) int {
	if seconds < 60 {
		fmt.Println("seconds cannot be below 60")
	}
	seconds /= 60
	return seconds
}

func fullName(firstName, lastName string) string {
	fullname := firstName + " " + lastName
	fmt.Sprint(fullname)
	return fullname

}

func divmod(a, b int) (int, int) {
	divide := a / b
	mod := a % b
	return divide, mod
}

func RectangleParams(height, width float64) (Area, Perimeter float64) {
	Area = height * width
	Perimeter = 2 * (height + width)
	return
}

func parseFullName(full string) (string, string) {
	sortName := strings.Fields(full)
	firstname := sortName[0]
	lastname := sortName[1]
	return firstname, lastname
}

func checkPassword(pass string) (bool, string) {
	if len(pass) >= 8 {
		return true, "pass"
	} else {
		return false, "the password is too short"
	}
}

func midpoint(x1, y1, x2, y2 float64) (float64, float64) {
	X := (x1 + x2) / 2
	Y := (y1 + y2) / 2
	return X, Y
}

func printSlice(nums []int) {
	fmt.Println(nums)
}

func sliceLength(s []string) int {
	return len(s)
}

func makefruits() []string {
	fruits := []string{"apple", "pear", "banana", "peach"}
	return fruits
}

func addToSlice(nums []int, value int) []int {
	nums = append(nums, value)
	return nums
}

func printPrices(prices map[string]float64) {
	fmt.Println(prices)
}

func makeScores() map[string]int {
	scores := map[string]int{
		"Alya": 10,
		"Mary": 8,
		"Jake": 5,
	}
	return scores
}

func applyDiscount(prices []float64, percent float64) []float64 {
	newSlice := []float64{}
	for _, price := range prices {
		percent /= 100
		price = price - (percent * price)
		newSlice = append(newSlice, price)
	}
	return newSlice
}

func wordFrequency(words []string) map[string]int {
	words = []string{"Python", "is", "a", "high-level,", "general-purpose", "programming", "language", "known", "for", "its", "clear,", "readable", "syntax", "and", "simplicity"}
	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++
	}

	return frequency
}

func filterEven(nums []int) []int {
	newNums := []int{}
	for _, num := range nums {
		if num%2 == 0 {
			newNums = append(newNums, num)
		}
	}
	return newNums
}

func mostExpensive(prices map[string]float64) (string, float64) {
	var (
		expensiveName  string
		expensivePrice float64
	)
	for name, price := range prices {
		if price > expensivePrice {
			expensivePrice = price
			expensiveName = name

		}
	}
	return expensiveName, expensivePrice
}

func averageGrade(grades map[string]int) float64 {
	grades = map[string]int{
		"Alya":   10,
		"Mary":   8,
		"Jake":   5,
		"Angela": 9,
		"Tom":    7,
	}
	sum := 0
	for _, grade := range grades {
		sum += grade
	}
	return float64(sum) / float64(len(grades))
}

func countLettersAndDigits(s string) (int, int) {
	letters := 0
	digits := 0

	for _, char := range s {
		if unicode.IsLetter(char) {
			letters++
		}

		if unicode.IsDigit(char) {
			digits++
		}
	}

	return letters, digits
}