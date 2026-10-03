package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

var n any = "Hello"

type (
	Vehicle interface {
		Start() string //task 1
		Stop() string
	}

	Shape interface {
		Area() float64 //task2
		Perimeter() float64
	}

	Employee interface {
		Salary() float64 //task3
		Position() string
	}

	Instrument interface { //task4
		Play() string
	}

	Warrior interface {
		Attack() int //task5
		Description() string
	}

	Transport interface {
		Price(distanceKm float64) float64 //task6
	}

	Discount interface {
		Apply(price float64) float64 //task7
	}

	Flyer interface {
		Fly() string
	}

	Swimmer interface {
		Swim() string
	}

	Amphibious interface {
		Flyer
		Swimmer
	}

	Depositor interface { //task15
		Deposit(amount float64) error
	}

	Withdrawer interface {
		Withdraw(amount float64) error
	}

	Account interface {
		Depositor
		Withdrawer
	}

	Handler interface { //task19
		Handle(event any)
	}

	Command interface {
		Execute() (string, error)
	}
)

var handlers = make(map[string][]Handler)

type (
	Car struct{ Name string } //task1

	Bike struct{ Name string }

	Bus struct{ Name string }

	Circle struct{ Radius float64 } //task2

	Rectangle struct {
		Height float64
		Width  float64
	}

	Triangle struct {
		A float64
		B float64
		C float64
	}

	Manager struct { //task3
		Name       string
		Basesalary float64
	}

	Developer struct {
		Name       string
		Basesalary float64
		Level      string
	}

	Guitar struct{ Name string } //task4

	Piano struct{ Name string }

	Drum struct{ Name string }

	Knight struct { //task5
		Name     string
		Strength int
	}

	Archer struct {
		Name      string
		Precision int
	}

	Mage struct {
		Name      string
		Manapower int
	}

	Taxi struct { //task6
		FixedKm    float64
		PricePerKm float64
	}

	Bus1 struct{ Price1 float64 }

	Bike1 struct{ PricePerKm float64 }

	PercentDiscound struct{ Percent float64 } //task7
	FixedDiscound   struct{ Amount float64 }
	NoDiscount      struct{}

	Point struct{ X, Y int } // task 8

	Box struct{ Value any } //task9

	Item struct {
		Name string
		Year int
	}

	Product struct {
		Name  string
		Price float64
	}

	Duck struct{ Name string }

	Wallet struct{ balance float64 } //task15

	Money struct { //task16
		Amount   float64
		Currency string
	}

	Task struct { //task17
		Title string
		Done  bool
	}

	PrinterHandler struct{} //task19

	CounterHandler struct{ counter int }

	StringOMGGETOUT struct{ Name string } //task21

	Division struct{ a, b int }

	Category struct {
		Name string
		Sub  []Category
	}

	Invoice struct {
		Client string
		Amount float64
	}

	Validation struct {
		Field   string
		Message string
	}
)

func (c Car) Start() string { return fmt.Sprintf("Car %v is strarting", c.Name) } //task1
func (c Car) Stop() string  { return fmt.Sprintf("Car %v stops", c.Name) }

func (b Bike) Start() string { return fmt.Sprintf("Bike %v is strarting", b.Name) }
func (b Bike) Stop() string  { return fmt.Sprintf("Bike %v stops", b.Name) }

func (b Bus) Start() string { return fmt.Sprintf("Bus %v is strarting", b.Name) }
func (b Bus) Stop() string  { return fmt.Sprintf("Bus %v stops", b.Name) }

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius } //task2
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2.0
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}
func (t Triangle) Perimeter() float64 { return t.A + t.B + t.C }

func TotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

func (m Manager) Salary() float64  { return m.Basesalary * 1.20 } //task3
func (m Manager) Position() string { return "Manager" }

func (d Developer) Salary() float64 {
	if d.Level == "senior" {
		return d.Basesalary * 1.30
	}
	return d.Basesalary
}
func (d Developer) Position() string { return "Developer (" + d.Level + ")" }

func TopGettingPaid(employee []Employee) float64 {
	topemployee := employee[0]

	for _, e := range employee[1:] {
		if e.Salary() > topemployee.Salary() {
			topemployee = e
		}
	}
	return topemployee.Salary()
}

func (g Guitar) Play() string { return fmt.Sprintln("Guitar" + g.Name + "plays: chika-chika") } //task4

func (p Piano) Play() string { return fmt.Sprintln("Piano" + p.Name + "plays: clunk-clink") }

func (d Drum) Play() string { return fmt.Sprintln("Drum" + d.Name + "plays: boom-boom") }

func (k Knight) Attack() int { return k.Strength * 2 } //task5
func (k Knight) Description() string {
	return fmt.Sprintf("%v's strength is %v. He attacks equal to %v", k.Name, k.Strength, k.Attack())
}

func (a Archer) Attack() int { return a.Precision + 10 }
func (a Archer) Description() string {
	return fmt.Sprintf("%v's precision is %v. He attacks equal to %v", a.Name, a.Precision, a.Attack())
}

func (m Mage) Attack() int { return m.Manapower * 3 }
func (m Mage) Description() string {
	return fmt.Sprintf("%v's mana is %v. His mana is powered equal to %v", m.Name, m.Manapower, m.Attack())
}

func SummaryAttack(warriors []Warrior) int {
	sum := 0
	for _, w := range warriors {
		sum += w.Attack()
	}
	return sum
}

func (t Taxi) Price(distanceKm float64) float64 { return t.FixedKm + distanceKm*t.PricePerKm } //task6

func (b Bus1) Price(distanceKm float64) float64 { return b.Price1 }

func (b Bike1) Price(distanceKm float64) float64 { return distanceKm * b.PricePerKm }

func (p PercentDiscound) Apply(price float64) float64 { //task7
	return price - price*p.Percent/100
}

func (f FixedDiscound) Apply(price float64) float64 {
	return price - f.Amount
}

func (n NoDiscount) Apply(price float64) float64 { return price }

func Describe(v any) { //task8
	fmt.Printf("Variable: %v, type: %T", v, v) //task8
}

func AsString(v any) (string, bool) { //task10
	if s, ok := v.(string); ok {
		return s, true
	}
	return "", false
}

func Classsify(v any) string { //task12
	switch x := v.(type) {
	case int:
		return fmt.Sprintln("integer")
	case float64:
		return fmt.Sprintln("float")
	case string:
		return fmt.Sprintln("string with length of", len(x))
	case bool:
		return fmt.Sprintln("boolean value")
	case nil:
		return fmt.Sprintln("empty")
	default:
		return fmt.Sprintf("Unknown type: %T", x)
	}
}

func Process(v any) { //task13
	switch x := v.(type) {
	case Product:
		fmt.Printf("Name: %v\nPrice: %.2f", x.Name, x.Price)
	case Shape:
		fmt.Println(x.Area())
	default:
		fmt.Println("uhh idk what to do with this, so uhh srry :( ")
	}
}

func (d Duck) Fly() string  { return fmt.Sprintln(d.Name + "flies!") } //task14
func (d Duck) Swim() string { return fmt.Sprintln(d.Name + "swims!") }

func (w *Wallet) Deposit(amount float64) error { //task15
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	w.balance += amount

	return nil
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if amount > w.balance {
		return errors.New("insufficient balance")
	}

	w.balance -= amount

	return nil
}

func (m Money) String() string { return fmt.Sprintln(m.Amount, m.Currency) } //task16

func (t Task) String() string { //task17
	if t.Done == false {
		return fmt.Sprintln("[ ]")
	}

	return fmt.Sprintln("[x]")
}

func LogMessage(w io.Writer, msg string) { //task18
	fmt.Fprintln(w, msg)
}

func Subscribe(eventName string, handler Handler) { //task19
	handlers[eventName] = append(handlers[eventName], handler)
}

func Publish(eventName string, event any) {
	for _, handler := range handlers[eventName] {
		handler.Handle(event)
	}
}

func (p PrinterHandler) Handle(event any) {
	fmt.Println("Events gained:", event)
}

func (c *CounterHandler) Handle(event any) {
	c.counter++

	fmt.Println("Counter: events handled", c.counter, "times")
}

func (s StringOMGGETOUT) Execute() (string, error) {
	return "Hello, Go!", nil
}

func (d Division) Execute() (string, error) {
	if d.b == 0 {
		return "", errors.New("division by zero")
	}

	result := d.a / d.b

	return fmt.Sprintf("%.2f", result), nil
}

func PrintTree(v any, depth int) { //task22
	switch x := v.(type) {
	case Category:
		fmt.Println(strings.Repeat("  ", depth) + x.Name)

		for _, sub := range x.Sub {
			PrintTree(sub, depth+1)
		}

	case []Category:
		for _, category := range x {
			PrintTree(category, depth)
		}
	}
}

func main() {
	fmt.Println("====== First task ======")
	vehicles := []Vehicle{
		Car{"Toyota"},
		Car{"Volfsvagen"},
		Bike{"Giant"},
		Bike{"Trek"},
		Bus{"Yutong"},
		Bus{"Daimer"},
	}

	for _, vehicle := range vehicles {
		fmt.Println(vehicle.Start())
		fmt.Println(vehicle.Stop())
	}

	fmt.Println("====== Second task ======")
	shapes := []Shape{
		Circle{23.4},
		Rectangle{15.8, 23.0},
		Circle{58.9},
		Triangle{24.9, 24.2, 23.9},
		Rectangle{56.6, 61.3},
		Triangle{45.9, 46.0, 46.6},
	}

	fmt.Println("Total sum of all area:", TotalArea(shapes))

	fmt.Println("====== Third task ======")
	employees := []Employee{
		Manager{Name: "Alisher", Basesalary: 5000.0},
		Developer{Name: "Alex", Basesalary: 7000.0, Level: "senior"},
		Developer{Name: "Matt", Basesalary: 4000.0, Level: "basic"},
		Manager{Name: "Aisha", Basesalary: 8000.0},
	}
	TopGettingPaid(employees)

	fmt.Println("====== Fourth task ======")
	instruments := []Instrument{
		Guitar{"Fender"},
		Guitar{"Gibson"},
		Piano{"Yamaha"},
		Piano{""},
		Drum{"Pearl Masters"},
		Drum{"Tama"},
	}

	for _, in := range instruments {
		fmt.Println(in.Play())
	}

	fmt.Println("====== Fifth task ======")
	warriors := []Warrior{
		Knight{"Tea Knight Cookie", 222},
		Archer{"Wind Archer Cookie", 182},
		Mage{"Mala Sauce Cookie", 248},
		Knight{"Elder Faerie Cookie", 258},
		Archer{"Creme Brule Cookie", 757},
		Mage{"Schwarzwälder", 159},
	}

	for _, w := range warriors {
		w.Description()
	}
	fmt.Println("Total attack:", SummaryAttack(warriors))

	fmt.Println("====== Sixth task ======")
	distance := 10.0
	transports := []Transport{
		Taxi{FixedKm: 5, PricePerKm: 2},
		Bus1{Price1: 4},
		Bike1{PricePerKm: 5},
	}
	cheapest := transports[0]
	cheapestPrice := cheapest.Price(distance)

	for _, transport := range transports {
		price := transport.Price(distance)

		fmt.Printf("%T: %.2f\n", transport, price)

		if price < cheapestPrice {
			cheapest = transport
			cheapestPrice = price
		}
	}

	fmt.Printf("The cheapest: %T — %.2f\n", cheapest, cheapestPrice)

	fmt.Println("====== Seventh task ======")
	price := 100.0

	discounts := []Discount{
		PercentDiscound{Percent: 20},
		FixedDiscound{Amount: 10},
		NoDiscount{},
	}

	fmt.Printf("Price before: %.2f\n", price)

	for _, discount := range discounts {
		price = discount.Apply(price)

		fmt.Printf("After %T: %.2f\n", discount, price)
	}

	fmt.Printf("Total price: %.2f\n", price)

	fmt.Println("====== Eighth task ======")
	Describe(42)
	Describe("hello")
	Describe(89.6)
	Describe(true)
	Describe([]int{3, 5, 7, 5, 6, 0})
	Describe(Point{X: 23, Y: 77})

	fmt.Println("====== Ninth task ======")
	box := []Box{
		{Value: 43},
		{Value: "hello"},
		{Value: []int{2, 4, 6}},
		{Value: Item{Name: "Algebra", Year: 2001}},
	}

	for _, b := range box {
		fmt.Printf("Value: %v, type: %T\n", b, b)
	}

	fmt.Println("====== Tenth task ======")
	fmt.Println(AsString(42))
	fmt.Println(AsString(45.9))
	fmt.Println(AsString([]int{6, 7, 8, 9}))
	fmt.Println(AsString("Golang"))

	fmt.Println("====== Eleventh task ======")
	values := []any{
		Circle{234},
		"hi",
		456.89,
		Rectangle{190, 192},
		Triangle{23, 18, 19},
		"I'm a shape pls trust",
	}
	totalAr := 0.0
	for _, v := range values {
		if shape, ok := v.(Shape); ok {
			totalAr += shape.Area()
		}

	}
	fmt.Printf("Summary area of shapes: %.2f", totalAr)

	fmt.Println("====== Twelth task ======")
	classify := []any{
		"hello",
		34,
		3.14,
		true,
		nil,
		[]int{4, 5, 6, 7},
	}

	for _, c := range classify {
		fmt.Println(Classsify(c))
	}

	fmt.Println("====== Thirteenth task ======")
	randoms := []any{
		Product{Name: "Shampoo", Price: 34.0},
		32,
		Rectangle{Height: 32, Width: 34},
		"Golang pls no",
		34.9,
		Triangle{A: 33, B: 31, C: 34},
		Product{Name: "Soap", Price: 20.0},
		Circle{Radius: 78},
	}

	for _, r := range randoms {
		Process(r)
	}

	fmt.Println("====== Fourteeth task ======")
	var animal Amphibious = Duck{Name: "Donald"}

	fmt.Println(animal.Fly())
	fmt.Println(animal.Swim())

	fmt.Println("====== Fifteenth task ======")
	var account Account = &Wallet{}

	err := account.Deposit(100)

	if err != nil {
		fmt.Println("Error:", err)
	}

	err = account.Withdraw(30)

	if err != nil {
		fmt.Println("Error:", err)
	}

	err = account.Withdraw(100)

	if err != nil {
		fmt.Println("Error:", err)

	}
	fmt.Println("====== Sixteenth task ======")
	money := []Money{
		{Amount: 150, Currency: "USD"},
		{Amount: 200, Currency: "EUR"},
		{Amount: 500, Currency: "TJS"},
	}

	for _, m := range money {
		fmt.Println(m)

		fmt.Println("====== Seventeenth task ======")
		tasks := []Task{
			{Title: "Изучить Go", Done: true},
			{Title: "Сделать домашку", Done: false},
			{Title: "Повторить интерфейсы", Done: true},
			{Title: "Изучить указатели", Done: false},
			{Title: "Написать программу", Done: true},
			{Title: "Повторить Stringer", Done: false},
		}

		for _, task := range tasks {
			fmt.Println(task)
		}

		fmt.Println("====== Eighteenth task ======")
		LogMessage(os.Stdout, "Привет, консоль!")
		var builder strings.Builder
		LogMessage(&builder, "Привет, буфер!")
		fmt.Println("Содержимое билдера: ", builder.String())

		fmt.Println("====== Nineteenth task ======")
		printer := PrinterHandler{}
		counter := &CounterHandler{}

		Subscribe("message", printer)
		Subscribe("message", counter)

		fmt.Println("First Publish:")
		Publish("message", "Hello!")

		fmt.Println("\nSecond Publish:")
		Publish("message", 42)

		fmt.Println("\nThird Publish:")
		Publish("message", true)

		fmt.Println("====== Twenth task ======")
		shapes := []Shape{
			Circle{Radius: 3},
			Rectangle{Width: 4, Height: 5},
			Triangle{A: 5, B: 5, C: 6},
		}

		largest := shapes[0]
		largestArea := largest.Area()

		for _, shape := range shapes[1:] {
			area := shape.Area()

			if area > largestArea {
				largest = shape
				largestArea = area
			}
		}

		fmt.Printf("Maximum area: %.2f\n", largestArea)

		switch shape := largest.(type) {
		case Circle:
			fmt.Println("Shape: Circle")
			fmt.Println("Radius:", shape.Radius)

		case Rectangle:
			fmt.Println("Shape: Rectangle")
			fmt.Println("Width:", shape.Width)
			fmt.Println("Height:", shape.Height)

		case Triangle:
			fmt.Println("Shape: Triangle")
			fmt.Println("Sides:", shape.A, shape.B, shape.C)
		}

		fmt.Println("====== Twenty first task ======")
		commands := map[string]Command{
			"hello":      StringOMGGETOUT{},
			"divide":     Division{a: 10, b: 2},
			"bad-divide": Division{a: 10, b: 0},
		}
		names := []string{
			"hello",
			"divide",
			"bad-divide",
			"unknown",
		}

		for _, name := range names {
			command, ok := commands[name]

			if !ok {
				fmt.Println("Command not found:", name)
				continue
			}

			result, err := command.Execute()

			if err != nil {
				fmt.Println("Error!:", err)
				continue
			}

			fmt.Println(name, "→", result)
		}

		fmt.Println("====== Twenty second task ======")
		rootCategory := Category{
			Name: "Games",
			Sub: []Category{
				{
					Name: "Minecraft",
					Sub: []Category{
						{Name: "Mods"},
						{Name: "Servers"},
					},
				},
				{
					Name: "Zelda",
					Sub: []Category{
						{Name: "Games"},
						{Name: "Lore"},
					},
				},
			},
		}

		PrintTree(rootCategory, 0)

	}
}
