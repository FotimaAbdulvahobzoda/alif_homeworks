package main

import (
	"fmt"
)

func main() {
	fmt.Println("===== First task =====")

	var (
		a int        //0
		b int8       //0
		c int16      //0
		d int32      //0
		e int64      //0
		f uint       //0
		g uint8      //0
		h uint16     //0
		i uint32     //0
		j uint64     //0
		k float32    //0
		l float64    //0
		m bool       //false
		n rune       // ''
		o byte       // 0
		p string     // ""
		q complex64  //(0+0i)
		r complex128 //(0+0i)
	)
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)

	a = 20
	b = 2
	c = -767
	d = 5_435
	e = -32_982_300
	f = 340
	g = 32
	h = 5_467
	i = 43_500
	j = 543_896_500
	k = 678.900
	l = 345.7960000
	m = true
	n = 'D'
	o = 56
	p = "Daniel"
	q = 3.43 - 2i
	r = 10.8 + 4i

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)

	fmt.Println("===== Second task =====")
	var (
		num1 int
		num2 int
	)
	fmt.Print("First number: ")
	fmt.Scan(&num1)
	fmt.Print("Second number: ")
	fmt.Scan(&num2)
	var (
		Addition       = num1 + num2
		Substraction   = num1 - num2
		Multiplication = num1 * num2
		Division       = num1 / num2
		Modulo         = num1 % num2
	)
	fmt.Printf("Addition: %d\nSubstraction = %d\nMultiplication: %d\nDivision: %d\nModulo: %d\n", Addition, Substraction, Multiplication, Division, Modulo)

	fmt.Println("===== Third task =====")
	const (
		PI          = 3.14
		CompanyName = "Alif"
		CurrentYear = 2026
		DaysInWeek  = 7
	)
	fmt.Println(PI, CompanyName, CurrentYear, DaysInWeek)

	fmt.Println("===== Fourth task =====")
	var (
		name      string
		age       int
		height    int
		weight    int
		doyouwork bool
	)
	fmt.Print("Name: ")
	fmt.Scan(&name)
	fmt.Print("Age: ")
	fmt.Scan(&age)
	fmt.Print("Height: ")
	fmt.Scan(&height)
	fmt.Print("Weight: ")
	fmt.Scan(&weight)
	fmt.Print("Do you work? (true/false): ")
	fmt.Scan(&doyouwork)
	if doyouwork {
		fmt.Print("you work")

	} else {
		fmt.Print("you don't work")
	}

	fmt.Printf("So, your name is %s, you're %d years old, you're %d meters and %d kilograms. And %v\n", name, age, height, weight, doyouwork)

	fmt.Println("===== Fifth task =====")
	var (
		number1 int
		number2 int
	)
	fmt.Print("First number: ")
	fmt.Scan(&number1)
	fmt.Println("Second number: ")
	fmt.Scan(&number2)
	fmt.Printf("Addition: %d\nSubstraction = %d\nMultiplication: %d\nDivision: %d\nModulo: %d\n", number1-number2, number1+number2, number1*number2, number1/number2, number1%number2)

	fmt.Println("===== Sixth task =====")
	fmt.Println("== DEVPASSPORT ==")
	var (
		name1                  string = "Далер"
		surname                string = "Муродов"
		age1                   int    = 29
		Seniority              string = "Сеньор-разработчик"
		waige                  int    = 120_000
		favcodelanguage        string = "Python, Go, C++"
		comercialdevexperience int    = 6
		doesheknowDocker       bool   = true
		hisinitial             rune   = 'Д'
		favasciisymbol         byte   = 76
		hisporjects            uint   = 32
	)
	fmt.Printf("Name: %s\n, Surname: %s\n, Age: %d\n, Seniority: %s\n, Waige: %d\n, Favorite coding language: %s\n, Commercial Developer experience: %d\n, Does he know Docker: %v\n, His initial = %v\n, Favorite ASCII symbol: %v\n, Personal projects: %v\n", name1, surname, age1, Seniority, waige, favcodelanguage, comercialdevexperience, doesheknowDocker, hisinitial, favasciisymbol, hisporjects)

	fmt.Println("===== Seventh task =====")
	var (
		num3 int = 3
		num4 int = 12
		num5 int = 9
	)
	fmt.Printf("First: %d\nSecond: %d\nThird: %d\n", num3, num4, num5)
	fmt.Println("--------------------")
	num3 = 6
	num4 = -32
	num5 = 367

	fmt.Printf("First: %d\nSecond: %d\nThird: %d\n", num3, num4, num5)
	fmt.Println("--------------------")
	num3 = 18
	num4 = -123
	num5 = 3456

	fmt.Printf("First: %d\nSecond: %d\nThird: %d\n", num3, num4, num5)

	fmt.Println("===== Eighth task =====")
	var (
		yourname    string
		yoursurname string
		ag4         int
		birthyear   int
		city        string
		profession  string
	)
	fmt.Println("=== My visit card ===")
	fmt.Print("Name: ")
	fmt.Scan(&yourname)
	fmt.Print("Surname: ")
	fmt.Scan(&yoursurname)
	fmt.Print("Age: ")
	fmt.Scan(&ag4)
	fmt.Print("Birthyear: ")
	fmt.Scan(&birthyear)
	fmt.Print("City: ")
	fmt.Scan(&city)
	fmt.Print("Job: ")
	fmt.Scan(&profession)

}
