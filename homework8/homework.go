package main

import (
	"fmt"
	l "fmt"
)

func main() {
	l.Println("====== First task ======")
	book := Book{Title: "Harry Potter: The Order of the Phoenix", Author: "J.K.Rolling", Year: 2003, Price: 25.00}
	l.Println(book)

	l.Println("====== Second task ======")
	car := Car{}
	l.Println(car)

	l.Println("====== Third task ======")
	student := Student{Name: "Anton"}
	l.Println(student)

	l.Println("====== Fourth task ======")
	point := Point{X: 45, Y: 77}
	getname := getName(point)
	getvar, getvar1 := getVarX(point), getVarY(point)
	l.Printf("Name: %+v\n ", getname)
	l.Println("Variable: ", getvar, getvar1)

	l.Println("====== Fifth task ======")
	laptop := Laptop{
		Brand: "Macbook Neo",
		RAM:   8,
		SSD:   256,
		Price: 500.0,
	}
	l.Println("Brand: ", laptop.Brand)
	l.Println("RAM: ", laptop.RAM, "GB")
	l.Println("SSD: ", laptop.SSD, "GB")
	l.Println("Price: ", laptop.Price, "$")

	l.Println("====== Sixth task ======")
	movie1 := Movie{
		Title:    "A boy in the stripped pyjamas",
		Duration: 94,
		Rating:   7.7,
	}

	movie2 := Movie{
		Title:    "1+1",
		Duration: 112,
		Rating:   8.6,
	}

	movie3 := Movie{
		Title:    "Spider Man: Brand New Day",
		Duration: 150,
		Rating:   8.0,
	}
	l.Println("==== Names ====")
	l.Println(movie1.Title, movie2.Title, movie3.Title)
	l.Println("==== Durations ====")
	l.Println(movie1.Duration, "min", movie2.Duration, "min", movie3.Duration, "min")
	l.Println("==== Ratings ====")
	l.Println(movie1.Rating, "/10", movie2.Rating, "/10", movie3.Rating, "/10")

	l.Println("====== Seventh task ======")
	city1 := City{"New York", 8_584_629, false}
	city2 := City{Population: 1_274_774, Name: "Dushanbe", isCapital: true} //there's no difference?
	l.Println(city1, "and", city2)

	l.Println("====== Eighth task ======")
	owner := Account{
		Owner:   "Alisher",
		balance: 1000,
	}
	l.Println(owner)
	owner.balance += 250
	l.Println(owner)
	l.Println("====== Ninth task ======")
	temp := Temperature{
		Celsius: 36.6,
	}

	F := temp.Celsius*9/5 + 32
	l.Println("Celsius: ", temp.Celsius, "Farenheight: ", F)

	l.Println("====== Tenth task ======")
	product1 := Product{
		Name:     "Milk",
		Price:    10.0,
		Quantity: 25,
	}

	product2 := Product{
		Name:     "Cheese",
		Price:    20.5,
		Quantity: 19,
	}
	product3 := Product{
		Name:     "Yogurt",
		Price:    9.0,
		Quantity: 20,
	}
	l.Println("Total cost of milk: ", total(product1))
	l.Println("Total cost of cheese: ", total(product2))
	l.Println("Total cost of yogurt: ", total(product3))

	l.Println("====== Eleventh task ======")
	person := Person{
		Name: "Fotima",
		Age:  17,
	}
	l.Println(person)
	if person.Age >= 18 {
		fmt.Println(person.Name, "is an adult")
	}
	fmt.Println(person.Name, "is a minor")

	l.Println("====== Twelth task ======")
	rect := Rectangle{
		Width:  30.7,
		Height: 50.3,
	}
	l.Println(rect)
	area := func(Width float64, Height float64) float64 {
		return Width + Height
	}
	ar := area(rect.Width, rect.Height)
	l.Println("Area: ", ar)

	parameter := func(a float64, b float64) float64 {
		return (a + b) * 2
	}
	par := parameter(rect.Width, rect.Height)
	l.Println("Parameter: ", par)

	l.Println("====== Thirteenth task ======")
	employee := Employee{
		Name:   "Ahmad",
		Salary: 10000,
	}
	l.Println("Salary before: ", employee.Salary)
	employee.Salary *= 1.15
	l.Println("Salary after: ", employee.Salary)

	l.Println("====== Fourteenth task ======")
	counter := Counter{
		Value: 0,
	}
	counter.Value += 1
	l.Println(counter)
	counter.Value += 1
	l.Println(counter)
	counter.Value += 1
	l.Println(counter)

	l.Println("====== Fifteenth task ======")

	person1 := Person1{
		Name: "Anisa",
		Age:  28,
		Address: Address{
			City:   "Dushanbe",
			Street: "Rudaki 98",
		},
	}
	l.Println(person1)
	l.Println(person1.Name, "lives in", person1.Address.City)

	l.Println("====== Sixteenth task ======")
	l.Println(person1)
	person1.Address.Street = "Tursunzoda 25"
	l.Println(person1)

	l.Println("====== Seventeeth task ======")
	engine := Car1{
		Brand: "Lamborgini Urus",
		Engine: Engine{
			HorsePower: 812,
			Fuel:       "95",
		},
	}
	l.Println(engine)
	l.Println("Horse power of ", engine.Brand, ":", engine.Engine.HorsePower)

	l.Println("====== Eighteenth task ======")
	box := Box{
		Name: "simple box",
		Dimensions: Dimensions{
			Width:  20.6,
			Height: 23.2,
			Depth:  30.0,
		},
	}
	l.Println(box)
	total := box.Dimensions.Width * box.Dimensions.Height * box.Dimensions.Depth
	l.Println(total)

	l.Println("====== Nineteeth task ======")
	citizen := Citizen{
		Name:     "Matthew",
		Passport: Passport{},
	}
	l.Println("Cittzen: ", citizen) //Cittzen:  {Matthew { }}

	l.Println("====== Twenth task ======")
	employee1 := Employee1{
		Name: "Greg",
		Company: Company{
			Name: "Amazon Company",
			Address: Address1{
				City:   "New York",
				Street: "Arizona 1/2",
				Zip:    "231231",
			},
		},
	}
	l.Println(employee1)
	l.Println("The employee works in", employee1.Company.Name, "in", employee1.Company.Address.City)

	l.Println("====== Twenty first task ======")
	car1 := Car2{
		Brand: "Lamborgini Urus",
		Engine: Engine1{
			Power: 812,
			Type:  "Gasoline",
		},
	}

	car2 := Car2{
		Brand: "BMW M5",
		Engine: Engine1{
			Power: 717,
			Type:  "Petrol",
		},
	}
	l.Println(car1)
	l.Println(car2)

	if car1.Engine.Power > car2.Engine.Power {
		fmt.Println(car1.Brand, "is more powerful than", car2.Brand)
	} else {
		l.Println(car2.Brand, "is more powerful than", car1.Brand)
	}

	l.Println("====== Twenty second task ======")
	car3 := Car3{
		Model: "BMV M5",
		FrontWheel: Wheel{
			Size:  21,
			Brand: "BMW M Perfomance Wheels",
		},
		RearWheel: Wheel{
			Size:  20,
			Brand: "BMW M Perfomance Wheels",
		},
	}

	l.Println(car3)

	if car3.FrontWheel.Size > car3.RearWheel.Size {
		fmt.Println("Front wheels are larger than rear wheels of", car3.Model)
	} else {
		fmt.Println("Rear wheels are larger than front wheels of", car3.Model)
	}

	l.Println("====== Twenty third task ======")
	books := []Book1{
		{Title: "Harry Potter: The Order of the Phoenix", Pages: 870, Price: 25.0},
		{Title: "Diary of Whimpy Kid", Pages: 70, Price: 9.99},
		{Title: "The Little Prince", Pages: 256, Price: 16.99},
		{Title: "No Longer Human", Pages: 176, Price: 15.00},
	}
	max := 0
	name := ""
	for p := range books {
		if books[p].Pages > max {
			max = books[p].Pages
			name = books[p].Title
		}

	}
	l.Println("Name:", name, " | Pages:", max)

	l.Println("====== Twenty fourth task ======")
	product := []Product1{
		{Name: "Milk", Price: 10.0},
		{Name: "Cheese", Price: 23.5},
		{Name: "Yogurt", Price: 15.5},
		{Name: "Butter", Price: 29.0},
		{Name: "Cottage cheese", Price: 14.0},
	}
	sum := 0.0

	for _, p := range product {
		sum += p.Price
	}
	l.Println("Total:", sum, "Average cost:", sum/float64(len(product)))

	l.Println("====== Twenty fifth task ======")
	order := []Order{
		{ID: 8793947, Amount: 3, Completed: true},
		{ID: 7839447, Amount: 4, Completed: false},
		{ID: 2228394, Amount: 7, Completed: true},
		{ID: 9906574, Amount: 1, Completed: false},
		{ID: 5729347, Amount: 10, Completed: true},
		{ID: 3898403, Amount: 8, Completed: true},
	}

	c := 0
	s := 0.0
	for l, i := range order {
		if order[l].Completed {
			c += 1
			s += i.Amount
		}
	}
	l.Println("Completed:", c, "Amount summary:", s)

	l.Println("====== Twenty sixth task ======")
	students := []Student1{
		{Name: "Alan", Grade: 100},
		{Name: "Katie", Grade: 30},
		{Name: "John", Grade: 79},
		{Name: "Olivia", Grade: 90},
		{Name: "Angela", Grade: 60},
		{Name: "Jacob", Grade: 30},
		{Name: "Elliot", Grade: 63},
		{Name: "Helen", Grade: 83},
	}

	count := map[string]int{"A": 0, "B": 0, "C": 0, "F": 0}
	var letterG string

	for _, s := range students {
		switch {
		case s.Grade >= 90:
			letterG = "A"

		case s.Grade <= 89 || s.Grade >= 75:
			letterG = "B"
		case s.Grade <= 74 || s.Grade >= 60:
			letterG = "C"
		case s.Grade < 60:
			letterG = "F"
		}

		count[letterG]++

	}

	l.Println(students, count)

}

type (
	Book struct {
		Title  string
		Author string
		Year   int
		Price  float64
	}

	Car struct {
		Brand string
		Model string
		Year  int
	}

	Student struct {
		Name   string
		Grade  int
		Passed bool
	}

	Point struct {
		X int
		Y int
	}
)

func getName(p Point) Point {
	return p
}

func getVarX(p Point) int {
	return p.X
}
func getVarY(p Point) int {
	return p.Y
}

type (
	Laptop struct {
		Brand string
		RAM   int
		SSD   int
		Price float64
	}

	Movie struct {
		Title    string
		Duration int
		Rating   float64
	}

	City struct {
		Name       string
		Population int
		isCapital  bool
	}

	Account struct {
		Owner   string
		balance float64
	}

	Temperature struct {
		Celsius float64
	}

	Product struct {
		Name     string
		Price    float64
		Quantity int
	}
)

func total(p Product) float64 {
	return p.Price * float64(p.Quantity)
}

type (
	Person struct {
		Name string
		Age  int
	}

	Rectangle struct {
		Width  float64
		Height float64
	}

	Employee struct {
		Name   string
		Salary float64
	}

	Counter struct {
		Value int
	}

	Address struct {
		City   string
		Street string
	}

	Person1 struct {
		Name    string
		Age     int
		Address Address
	}

	Engine struct {
		HorsePower int
		Fuel       string
	}

	Car1 struct {
		Brand  string
		Engine Engine
	}

	Dimensions struct {
		Width  float64
		Height float64
		Depth  float64
	}

	Box struct {
		Name       string
		Dimensions Dimensions
	}

	Passport struct {
		Number   string
		Issuedby string
	}

	Citizen struct {
		Name     string
		Passport Passport
	}

	Address1 struct {
		City   string
		Street string
		Zip    string
	}

	Company struct {
		Name    string
		Address Address1
	}

	Employee1 struct {
		Name    string
		Company Company
	}

	Engine1 struct {
		Power int
		Type  string
	}

	Car2 struct {
		Brand  string
		Engine Engine1
	}

	Wheel struct {
		Size  int
		Brand string
	}

	Car3 struct {
		Model      string
		FrontWheel Wheel
		RearWheel  Wheel
	}

	Book1 struct {
		Title string
		Pages int
		Price float64
	}

	Product1 struct {
		Name  string
		Price float64
	}

	Order struct {
		ID        int
		Amount    float64
		Completed bool
	}

	Student1 struct {
		Name  string
		Grade int
	}
)
