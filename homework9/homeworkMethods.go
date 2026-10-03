package main

import (
	"errors"
	"fmt"
	l "fmt"
	"math"
	"strconv"
)

type Wallet struct {
	balance float64
}

func (w *Wallet) Deposit(amount float64) error {
	if amount < 0 {
		return l.Errorf("amount shouldn't be negative")
	}
	w.balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount < 0 {
		return l.Errorf("amount shouldn't be negative")
	}
	if amount > w.balance {
		return l.Errorf("not enough recources. You have %.2f, you need %.2f", w.balance, amount)
	}
	w.balance -= amount
	return nil
}

func (w Wallet) Balance() float64 {
	return w.balance
}

type (
	Song struct {
		Title    string
		Duration int
	}

	Playlist struct {
		song []Song
	}
)

func (p *Playlist) Add(s Song) {
	p.song = append(p.song, s)
}

func (p *Playlist) TotalDuration() int {

	total := 0

	for _, play := range p.song {
		total += play.Duration
	}

	return total
}

func (p *Playlist) Longest() Song {
	if len(p.song) == 0 {
		return Song{}
	}
	longest := p.song[0]
	for _, p := range p.song[1:] {
		if p.Duration > longest.Duration {
			longest = p
		}
	}
	return longest
}

func (p *Playlist) String() string {
	res := ""
	for i, s := range p.song {
		res = l.Sprintf("%d. %s: %d", i+1, s.Title, s.Duration)
	}
	return res
}

type Inventory struct {
	items map[string]int
}

func (inv *Inventory) Add(item string, qty int) {
	inv.items[item] += qty

}

func (inv *Inventory) Remove(item string, qty int) error {
	count, ok := inv.items[item]

	if !ok {
		return errors.New("item not found")
	}
	if count < qty {
		return errors.New("not enough resources")
	}
	inv.items[item] -= qty
	return nil
}

func (inv *Inventory) Count(item string, qty int) int {
	return inv.items[item]
}

func (inv *Inventory) Total() int {
	sum := 0
	for _, q := range inv.items {
		sum += q
	}
	return sum
}

type Matrix struct {
	mat [][]float64
}

func (m Matrix) Rows() int {
	return len(m.mat)
}

func (m Matrix) Cols() int {
	if len(m.mat) == 0 {
		return 0
	}
	return len(m.mat[0])
}

func (m Matrix) Total() float64 {
	sum := 0.0
	for _, row := range m.mat {
		for _, val := range row {
			sum += val
		}
	}
	return sum
}

func (m Matrix) Transpose() Matrix {
	rows := m.Rows()
	cols := m.Cols()

	resultData := make([][]float64, cols)
	for i := range resultData {
		resultData[i] = make([]float64, rows)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			resultData[c][r] = m.mat[r][c]
		}
	}

	return Matrix{mat: resultData}
}

type Point struct {
	X, Y float64
}

type Polygon struct {
	Points []Point
}

func (p Point) DistanceTo(other Point) float64 {
	dX := p.X - other.X
	dY := p.Y - other.Y

	return math.Sqrt(dX*dX + dY*dY)
}

func (p Polygon) Perimeter() float64 {
	total := 0.0

	for i := 0; i < len(p.Points); i++ {
		current := p.Points[i]
		next := p.Points[(i+1)%len(p.Points)]

		total = current.DistanceTo(next)
	}
	return total
}

type Employee struct {
	Name       string
	Salary     float64
	YearWorked int
}

func (e Employee) Bonus() float64 {
	periods := e.YearWorked / 5
	bonus := float64(periods) * 0.10
	return e.Salary * bonus
}

type Company struct {
	Employee []Employee
}

func (c Company) TotalPayRoll() float64 {
	total := 0.0

	for _, p := range c.Employee {
		total += p.Salary + p.Bonus()
	}
	return total
}

func (c Company) TopEarner() Employee {
	top := c.Employee[0]

	for _, employee := range c.Employee[1:] {
		currentTotal := employee.Salary + employee.Bonus()
		topTotal := top.Salary + top.Bonus()

		if currentTotal > topTotal {
			top = employee
		}
	}

	return top
}

type Order struct {
	ID     int
	Amount float64
}

type OrderQueue struct {
	orders []Order
}

func (q *OrderQueue) Enqueue(o Order) {
	q.orders = append(q.orders, o)
}

func (q *OrderQueue) Dequeue() (Order, error) {
	if len(q.orders) == 0 {
		return Order{}, errors.New("queue is empty")
	}

	first := q.orders[0]
	q.orders = q.orders[1:]

	return first, nil
}

func (q *OrderQueue) IsEmpty() bool {
	return len(q.orders) == 0
}

func (q *OrderQueue) TotalAmount() float64 {
	total := 0.0

	for _, order := range q.orders {
		total += order.Amount
	}

	return total
}

type Animal struct {
	Name  string
	Sound string
}

func (a Animal) MakeSound() string {
	return a.Name + " говорит " + a.Sound
}

type Dog struct {
	Animal
}

func (d Dog) MakeSound() string {
	return d.Animal.MakeSound() + " (виляет хвостом)"
}

type Cat struct {
	Animal
}

func (c Cat) MakeSound() string {
	return c.Animal.MakeSound() + " (мурлычет)"
}

type Speaker interface {
	MakeSound() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Info() string {
	return p.Name + ", " + strconv.Itoa(p.Age) + " лет"
}

type Student struct {
	Person
	University string
}

func (s Student) Info() string {
	return s.Person.Info() + ", университет: " + s.University
}

type Node struct {
	Value int
	Next  *Node
}

func (head *Node) Append(v int) *Node {
	newNode := &Node{
		Value: v,
	}

	if head == nil {
		return newNode
	}

	current := head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode

	return head
}

func PrintList(head *Node) {
	current := head

	for current != nil {
		l.Print(current.Value)

		if current.Next != nil {
			l.Print(" -> ")
		}

		current = current.Next
	}

	l.Println()
}

type Robot struct {
	x int
	y int
}

func (r *Robot) MoveUp() {
	r.y++
}

func (r *Robot) MoveDown() {
	r.y--
}

func (r *Robot) MoveLeft() {
	r.x--
}

func (r *Robot) MoveRight() {
	r.x++
}

func (r Robot) Position() (int, int) {
	return r.x, r.y
}

func (r Robot) String() string {
	return l.Sprintf("Robot @ (%d, %d)", r.x, r.y)
}

type Student1 struct {
	Name   string
	Grades []int
}

func (s Student1) Average() float64 {
	if len(s.Grades) == 0 {
		return 0
	}

	total := 0

	for _, grade := range s.Grades {
		total += grade
	}

	return float64(total) / float64(len(s.Grades))
}

func (s Student1) String() string {
	return l.Sprintf("%s: %.2f", s.Name, s.Average())
}

type BankAccount struct {
	Owner   string
	balance float64
}

func (a *BankAccount) Deposit(amount float64) {
	a.balance += amount
}

func (a *BankAccount) Withdraw(amount float64) error {
	if amount > a.balance {
		return errors.New("not enough money")
	}

	a.balance -= amount

	return nil
}

func Transfer(from, to *BankAccount, amount float64) error {
	err := from.Withdraw(amount)

	if err != nil {
		return err
	}

	to.Deposit(amount)

	return nil
}

type Task struct {
	Title    string
	Priority int
	Done     bool
}

func (t *Task) Complete() {
	t.Done = true
}

type User struct {
	ID   int
	Name string
}

func NewIDGenerator() func() int {
	nextID := 1

	return func() int {
		id := nextID
		nextID++

		return id
	}
}

type Category struct {
	Name string
	Sub  []Category
}

func CountAll(c Category) int {
	total := 1

	for _, sub := range c.Sub {
		total += CountAll(sub)
	}

	return total
}

type Contact struct {
	Name  string
	Phone string
}

type AddressBook struct {
	contacts map[string]Contact
}

func NewAddressBook() *AddressBook {
	return &AddressBook{
		contacts: make(map[string]Contact),
	}
}

func (a *AddressBook) Add(key string, c Contact) {
	a.contacts[key] = c
}

func (a *AddressBook) Find(key string) (Contact, bool) {
	contact, ok := a.contacts[key]
	return contact, ok
}

func (a *AddressBook) Delete(key string) {
	delete(a.contacts, key)
}

type Character struct {
	Name   string
	HP     int
	Attack int
}

func (c Character) IsAlive() bool {
	return c.HP > 0
}

func (c *Character) TakeDamage(dmg int) {
	c.HP -= dmg

	if c.HP < 0 {
		c.HP = 0
	}
}

func (c *Character) AttackTarget(target *Character) {
	target.TakeDamage(c.Attack)
}

type TrafficLight struct {
	color string
}

func NewTrafficLight() *TrafficLight {
	return &TrafficLight{
		color: "red",
	}
}

func (t *TrafficLight) Next() {
	switch t.color {
	case "red":
		t.color = "yellow"

	case "yellow":
		t.color = "green"

	case "green":
		t.color = "red"
	}
}

func (t TrafficLight) Current() string {
	return t.color
}

func (t TrafficLight) String() string {
	return "Traffic light: " + t.color
}

type Stack struct {
	values []int
}

func (s *Stack) Push(v int) {
	s.values = append(s.values, v)
}

func (s *Stack) Pop() (int, error) {
	if len(s.values) == 0 {
		return 0, errors.New("stack is empty")
	}

	lastIndex := len(s.values) - 1
	value := s.values[lastIndex]

	s.values = s.values[:lastIndex]

	return value, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}

type Player struct {
	Name  string
	Score int
}

func (p *Player) AddPoints(points int) {
	p.Score += points
}

func TopPlayer(players []Player) Player {
	top := players[0]

	for _, player := range players[1:] {
		if player.Score > top.Score {
			top = player
		}
	}

	return top
}

type Config struct {
	Host    string
	Port    int
	Timeout int
}

func NewConfig(host string, port int) *Config {
	return &Config{
		Host:    host,
		Port:    port,
		Timeout: 30,
	}
}

type Canvas struct {
	pixels [][]rune
}

func NewCanvas(width, height int) *Canvas {
	pixels := make([][]rune, height)

	for y := 0; y < height; y++ {
		pixels[y] = make([]rune, width)

		for x := 0; x < width; x++ {
			pixels[y][x] = ' '
		}
	}

	return &Canvas{
		pixels: pixels,
	}
}

func (c *Canvas) SetPixel(x, y int, char rune) error {
	if y < 0 || y >= len(c.pixels) ||
		x < 0 || x >= len(c.pixels[0]) {
		return errors.New("coordinates out of bounds")
	}

	c.pixels[y][x] = char

	return nil
}

func (c Canvas) Print() {
	for _, row := range c.pixels {
		l.Println(string(row))
	}
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func NewRectangle(width, height float64) (*Rectangle, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("width and height must be positive")
	}

	return &Rectangle{
		width:  width,
		height: height,
	}, nil
}

type Queue struct {
	values []int
}

func NewQueue() *Queue {
	return &Queue{
		values: []int{},
	}
}

func (q *Queue) Enqueue(v int) {
	q.values = append(q.values, v)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.values) == 0 {
		return 0, errors.New("queue is empty")
	}

	value := q.values[0]
	q.values = q.values[1:]

	return value, nil
}

type Car struct {
	Brand   string
	Model   string
	mileage int
}

func (c *Car) Drive(km int) {
	c.mileage += km
}

func NewCar(brand, model string) *Car {
	return &Car{
		Brand:   brand,
		Model:   model,
		mileage: 0,
	}
}

func NewUsedCar(brand, model string, mileage int) *Car {
	return &Car{
		Brand:   brand,
		Model:   model,
		mileage: mileage,
	}
}

type Logger struct {
	prefix string
	count  int
}

func NewLogger(prefix string) *Logger {
	return &Logger{
		prefix: prefix,
		count:  0,
	}
}

func (l *Logger) Log(msg string) {
	l.count++

	fmt.Printf("[%s] %s (#%d)\n", l.prefix, msg, l.count)
}

func main() {
	l.Println("====== First task ======")
	w := Wallet{100}
	w.Deposit(100)
	l.Printf("Balance: %.2f\n", w.Balance())
	w.Deposit(50)
	l.Printf("Balance: %.2f\n", w.Balance())
	if err := w.Withdraw(200); err != nil {
		l.Printf("Error: %v\n", err)

	}
	w.Withdraw(30)
	l.Printf("Balance: %.2f\n", w.Balance())

	l.Println("====== Second task ======")
	p := Playlist{}
	p.Add(Song{"Song 1", 180})
	p.Add(Song{"Song 2", 240})
	p.Add(Song{"Song 3", 150})
	p.Add(Song{"Song 4", 310})
	l.Printf("List of songs:\n%v", p)
	l.Printf("Total duration: %d sec\n", p.TotalDuration())
	l.Printf("The longest song: %s (%d sec)\n\n", p.Longest().Title, p.Longest().Duration)

	l.Println("====== Third task ======")
	inv := Inventory{
		items: map[string]int{
			"Apples": 10,
		},
	}
	inv.Add("Apples", 10)
	inv.Add("Pears", 5)
	l.Printf("Apples in the storage: %d\n", inv.Count("Apples", 5))
	l.Printf("Total: %d\n", inv.Total())
	if err := inv.Remove("Яблоки", 15); err != nil {
		l.Printf("Ошибка списания: %v\n", err)
	}

	l.Println("====== Fourth task =====")
	m := Matrix{
		mat: [][]float64{
			{1.0, 2.0, 3.0},
			{4.0, 5.0, 6.0},
		},
	}
	l.Println("Rows:", m.Rows())
	l.Println("Columns:", m.Cols())
	l.Println("Sum of the elements: ", m.Total())

	transposed := m.Transpose()
	l.Println("\nTransposing the maxtrix (3x2):")
	for _, row := range transposed.mat {
		l.Println(row)
	}

	l.Println("====== Fifth task =====")
	poly := Polygon{
		Points: []Point{
			{X: 0, Y: 0},
			{X: 0, Y: 3},
			{X: 4, Y: 0},
		},
	}
	l.Printf("Периметр треугольника: %.2f\n", poly.Perimeter())

	//task 8
	l.Println("====== Eighth task ======")
	dog := Dog{
		Animal: Animal{
			Name:  "Bobik",
			Sound: "woof",
		},
	}

	cat := Cat{
		Animal: Animal{
			Name:  "Murka",
			Sound: "meow",
		},
	}

	animals := []Speaker{dog, cat}

	for _, animal := range animals {
		l.Println(animal.MakeSound())
	}

	//task 9
	l.Println("====== Ninth task ======")
	student := Student{
		Person: Person{
			Name: "Alex",
			Age:  18,
		},
		University: "MIT",
	}

	l.Println(student.Info())

	// task 10
	l.Println("====== Tenth task ======")
	head := &Node{Value: 10}

	head.Append(20)
	head.Append(30)
	head.Append(40)
	head.Append(50)

	PrintList(head)

	// task 11
	l.Println("====== Eleventh task ======")
	students := []Student1{
		{"Alex", []int{5, 4, 5}},
		{"Bob", []int{4, 4, 5}},
		{"Mike", []int{5, 5, 5}},
		{"John", []int{3, 5, 4}},
	}

	best := students[0]

	for _, student := range students[1:] {
		if student.Average() > best.Average() {
			best = student
		}
	}

	l.Println("The best student:", best)

	// task 12
	l.Println("====== Twelth task ======")
	alex := BankAccount{Owner: "Alex"}
	bob := BankAccount{Owner: "Bob"}

	alex.Deposit(1000)

	err := Transfer(&alex, &bob, 300)

	if err != nil {
		l.Println(err)
	} else {
		l.Println("Transfer successful")
	}

	l.Println(alex)
	l.Println(bob)

	err = Transfer(&alex, &bob, 1000)

	if err != nil {
		l.Println("Error:", err)
	}

	l.Println(alex)
	l.Println(bob)

	// task 13
	l.Println("====== Thirteenth task ======")
	tasks := []Task{
		{"Learn Go", 5, false},
		{"Do homework", 2, false},
		{"Play Minecraft", 8, false},
		{"Read book", 1, false},
		{"Practice coding", 4, false},
		{"Clean room", 3, false},
	}

	l.Println("Before sorting:")

	for _, task := range tasks {
		l.Println(task)
	}

	for i := 0; i < len(tasks)-1; i++ {
		for j := 0; j < len(tasks)-1-i; j++ {
			if tasks[j].Priority > tasks[j+1].Priority {
				tasks[j], tasks[j+1] = tasks[j+1], tasks[j]
			}
		}
	}

	l.Println("\nAfter sorting:")

	for _, task := range tasks {
		l.Println(task)
	}

	// task 14
	l.Println("====== Fourteenth task ======")
	generator := NewIDGenerator()

	users := []User{
		{ID: generator(), Name: "Alex"},
		{ID: generator(), Name: "Bob"},
		{ID: generator(), Name: "Mike"},
		{ID: generator(), Name: "John"},
	}

	for _, user := range users {
		l.Println(user)
	}

	// task 15
	l.Println("====== Fifteenth task ======")
	electronics := Category{
		Name: "Electronics",
		Sub: []Category{
			{
				Name: "Phones",
				Sub: []Category{
					{Name: "Android"},
					{Name: "iPhone"},
				},
			},
			{
				Name: "Computers",
			},
		},
	}

	l.Println(CountAll(electronics))

	// task 16
	l.Println("====== Sixteenth task ======")
	book := NewAddressBook()

	book.Add("alex", Contact{
		Name:  "Alex",
		Phone: "12345",
	})

	book.Add("bob", Contact{
		Name:  "Bob",
		Phone: "67890",
	})

	contact, ok := book.Find("alex")

	if ok {
		l.Println("Found:", contact)
	} else {
		l.Println("Contact not found")
	}

	book.Delete("alex")

	contact, ok = book.Find("alex")

	if ok {
		l.Println("Found:", contact)
	} else {
		l.Println("Contact not found")
	}

	// Несуществующий контакт
	contact, ok = book.Find("mike")

	if ok {
		l.Println("Found:", contact)
	} else {
		l.Println("Contact not found")
	}

	//task 17
	l.Println("====== Seventeenth task ======")
	kazuya := Character{
		Name:   "Kazuya",
		HP:     100,
		Attack: 25,
	}

	jin := Character{
		Name:   "Jin",
		HP:     100,
		Attack: 20,
	}

	for kazuya.IsAlive() && jin.IsAlive() {
		kazuya.AttackTarget(&jin)

		l.Printf(
			"%s attacks %s → %s HP: %d\n",
			kazuya.Name,
			jin.Name,
			jin.Name,
			jin.HP,
		)

		if !jin.IsAlive() {
			break
		}

		jin.AttackTarget(&kazuya)

		l.Printf(
			"%s attacks %s → %s HP: %d\n",
			jin.Name,
			kazuya.Name,
			kazuya.Name,
			kazuya.HP,
		)
	}

	if kazuya.IsAlive() {
		l.Println("Winner:", kazuya.Name)
	} else {
		l.Println("Winner:", jin.Name)
	}

	// task 18
	l.Println("====== Eighteenth task ======")
	robot := Robot{}

	l.Println(robot)

	robot.MoveUp()
	l.Println(robot)

	robot.MoveRight()
	l.Println(robot)

	robot.MoveRight()
	l.Println(robot)

	robot.MoveDown()
	l.Println(robot)

	robot.MoveLeft()
	l.Println(robot)

	x, y := robot.Position()

	l.Println("X:", x)
	l.Println("Y:", y)

	// task 19
	l.Println("====== Nineteenth task ======")
	light := NewTrafficLight()

	l.Println(light)

	for i := 0; i < 7; i++ {
		light.Next()
		l.Println(light)
	}

	//task 20
	l.Println("====== Twenth task ======")
	var stack Stack

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)

	value, err := stack.Pop()

	if err != nil {
		l.Println(err)
	} else {
		l.Println("Pop:", value)
	}

	value, err = stack.Pop()

	if err != nil {
		l.Println(err)
	} else {
		l.Println("Pop:", value)
	}

	l.Println("Empty:", stack.IsEmpty())

	stack.Pop()
	stack.Pop()

	_, err = stack.Pop()

	if err != nil {
		l.Println("Error:", err)
	}

	// task 21
	l.Println("====== Twenty first task ======")
	players := []Player{
		{"Alex", 0},
		{"Bob", 0},
		{"Mike", 0},
		{"John", 0},
	}

	players[0].AddPoints(50)
	players[1].AddPoints(100)
	players[2].AddPoints(70)
	players[3].AddPoints(80)

	players[0].AddPoints(40)
	players[2].AddPoints(50)

	for _, player := range players {
		l.Println(player)
	}

	top := TopPlayer(players)

	l.Println("Winner:", top.Name)
	l.Println("Score:", top.Score)

	// task 22
	l.Println("====== Twenty second task ======")
	canvas := NewCanvas(8, 4)

	canvas.SetPixel(1, 1, '*')
	canvas.SetPixel(2, 1, '*')
	canvas.SetPixel(3, 1, '*')
	canvas.SetPixel(4, 1, '*')
	canvas.SetPixel(5, 1, '*')

	canvas.Print()

	err = canvas.SetPixel(10, 10, '#')

	if err != nil {
		l.Println("Error:", err)
	}

	// task23
	l.Println("====== Twenty third task ======")
	config := NewConfig("localhost", 8080)

	l.Println("Host:", config.Host)
	l.Println("Port:", config.Port)
	l.Println("Timeout:", config.Timeout)

	// task 24
	l.Println("====== Twenty fourth task ======")
	rectangle, err := NewRectangle(-5, 10)

	if err != nil {
		l.Println("Error:", err)
	} else {
		l.Println("Area:", rectangle.Area())
		l.Println("Perimeter:", rectangle.Perimeter())
	}

	rectangle, err = NewRectangle(5, 10)

	if err != nil {
		l.Println("Error:", err)
	} else {
		l.Println("Area:", rectangle.Area())
		l.Println("Perimeter:", rectangle.Perimeter())
	}

	// task 25
	l.Println("====== Twenty fifth task ======")
	queue := NewQueue()

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	value, _ = queue.Dequeue()
	l.Println(value)

	value, _ = queue.Dequeue()
	l.Println(value)

	value, _ = queue.Dequeue()
	l.Println(value)

	_, err = queue.Dequeue()

	if err != nil {
		l.Println("Error:", err)
	}

	// task 26
	l.Println("====== Twenty sixth task ======")
	newCar := NewCar("Toyota", "Camry")
	usedCar := NewUsedCar("Honda", "Civic", 50000)

	newCar.Drive(100)
	usedCar.Drive(200)

	l.Println(newCar)
	l.Println(usedCar)

	// task 27
	l.Println("====== Twenty seventh task ======")
	serverLogger := NewLogger("SERVER")
	errorLogger := NewLogger("ERROR")

	serverLogger.Log("Server started")
	serverLogger.Log("Player connected")
	serverLogger.Log("Server stopped")

	errorLogger.Log("Something went wrong")
	errorLogger.Log("Connection failed")

}
