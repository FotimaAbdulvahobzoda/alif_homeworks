package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// =================================== TASK 1 ==============================================
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("You can't extract negative number out of square root")
	}
	math.Sqrt(x)
	return x, nil
}

// =================================== TASK 2 ==============================================
func Divide(a, b float64) (float64, error) {
	if a == 0 || b == 0 {
		return 0, fmt.Errorf("You can't divide %v by %v", a, b)
	}
	return a / b, nil
}

// =================================== TASK 3 ==============================================
func Withdraw(balance, amount float64) (float64, error) {
	if amount < 0 {
		return 0, fmt.Errorf("Amount cannot be negative")
	} else if amount > balance {
		return 0, fmt.Errorf("unsifficient amount")
	}

	balance -= amount
	return balance, nil
}

// =================================== TASK 4 ==============================================
func step1(x int) error {
	if x%2 == 0 {
		return errors.New("this number is divided by two!")
	}
	return nil
}

func step2(x int) error {
	if x%3 == 0 {
		return errors.New("this number is divided by three!")
	}
	return nil
}

func step3(x int) error {
	if x%4 == 0 {
		return errors.New("this number is divided by four!")
	}
	return nil
}

func Pipeline(n int) error {
	err := step1(n)
	err2 := step2(n)
	err3 := step3(n)
	if err != nil {
		return fmt.Errorf("Error!: %w", err)
	} else if err2 != nil {
		return fmt.Errorf("Error!: %w", err2)
	} else if err3 != nil {
		return fmt.Errorf("Error!: %w", err3)
	}
	return nil
}

var (
	ErrNegative      = errors.New("age cannot be negative")
	ErrTooOld        = errors.New("age cannot be THAT old")
	ErrEmptyLogin    = errors.New("login shouldn't be empty. Try again!")
	ErrShortPassword = errors.New("the password is too short. Try again!")
	ErrNotFound      = errors.New("item not found")
	ErrUserExits     = errors.New("this user already exists")
)

// =================================== TASK 5 ==============================================
func Validage(age int) error {
	if age < 0 {
		return ErrNegative
	} else if age > 120 {
		return ErrTooOld
	}
	return nil
}

// =================================== TASK 6 ==============================================
func CheckCredentials(login, password string) error {
	if login == "" {
		return ErrEmptyLogin
	} else if len(password) < 6 {
		return ErrShortPassword
	}
	return nil
}

// =================================== TASK 7 ==============================================
func findItem(id int) error {
	if id == 0 {
		return ErrNotFound
	}
	return nil
}

func loadItem(id int) error {
	err := findItem(id)
	if err != nil {
		return fmt.Errorf("Error! %w. Couldn't load items %d", err, id)
	}
	return nil
}

// =================================== TASK 8 ==============================================
func readConfig() error {
	return errors.New("something went wrong")
}

func Startserver() error {
	return fmt.Errorf("%w. And we don't know what", readConfig())
}

func run() error {
	err := Startserver()

	if err != nil {
		return fmt.Errorf("and we're so worried: %w", err)
	}

	return nil
}

// =================================== TASK 9 ==============================================
func ParseAge(s string) (int, error) {
	res, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("%w. couldn't extract age", err)
	}
	return res, nil
}

func saveDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Division failed")
		}
	}()
	result = a / b
	return
}

func ValidatePassword(password string) []error {
	var errs []error
	if len(password) < 8 {
		errs = append(errs, errors.New("the password is too short"))
	}

	hasDigit := false
	hasUpper := false
	for _, char := range password {
		if char >= '0' && char <= '9' {
			hasDigit = true
		}
		if char >= 'A' && char <= 'Z' {
			hasUpper = true
		}
	}
	if !hasDigit {
		errs = append(errs, errors.New("the password should contain at least a number"))
	}
	if !hasUpper {
		errs = append(errs, errors.New("the password should contain at least an Upper word"))
	}
	return errs
}

func Retry(attempts int, f func() error) error {
	var lastErr error

	for i := 0; i < attempts; i++ {
		err := f()

		if err == nil {
			return nil
		}

		lastErr = err
	}

	return fmt.Errorf("All attempts are gone: %w", lastErr)
}

func Register(login string, users map[string]bool) error {
	if users[login] {
		return ErrUserExits
	}
	users[login] = true
	return nil
}

func RegisterWithRetry(baseLogin string, users map[string]bool) string {
	login := baseLogin
	suffix := 2

	for {
		err := Register(login, users)

		if err == nil {
			return login
		}

		if errors.Is(err, ErrUserExits) {
			login = fmt.Sprintf("%s_%d", baseLogin, suffix)
			suffix++
			continue
		}

		return ""
	}
}

func ProcessBatch(items []string, process func(string) error) (successCount int, errs []error) {
	for _, item := range items {
		err := process(item)

		if err != nil {
			errs = append(errs, err)
			continue
		}

		successCount++
	}

	return
}

func main() {
	fmt.Println("====== First task ======")
	squares := []float64{3.45, -3, 23.56, 789.98, 0, -12.4}
	for _, s := range squares {
		res, err := Sqrt(s)
		if err != nil {
			fmt.Println("Error!:", err)
		} else {
			fmt.Println("The result:", res)
		}
	}

	fmt.Println("====== Second task ======")
	firstNumber := []float64{2.34, 34, 0, 1, 89, 4, 0, 0}
	secondNumber := []float64{2, 2.9, 10, 0, 0, 56, 99, 3}
	for _, f := range firstNumber {
		for _, s := range secondNumber {
			res, err := Divide(f, s)
			if err != nil {
				fmt.Println("Error!", err)
			} else {
				fmt.Printf("The result of %v by %v: %v\n", f, s, res)
			}
		}
	}

	fmt.Println("====== Third task ======")
	res, err := Withdraw(234, 78)
	if err != nil {
		fmt.Println("Error!:", err)
	} else {
		fmt.Println("Withdraw complete! Now you have", res, "left")
	}

	res1, err := Withdraw(90, 0)
	if err != nil {
		fmt.Println("Error!:", err)
	} else {
		fmt.Println("Withdraw complete! Now you have", res1, "left")
	}

	res2, err := Withdraw(400, 450)
	if err != nil {
		fmt.Println("Error!:", err)
	} else {
		fmt.Println("Withdraw complete! Now you have", res2, "left")
	}

	fmt.Println("====== Fourth task ======")
	nums := []int{2, 23, 44, 55, 66, 67, 82}
	for _, n := range nums {
		err := Pipeline(n)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("the numbers are valid")
		}
	}

	fmt.Println("====== Fifth task ======")
	ages := []int{-3, 23, 18, 56, 156, -99, 110}

	for _, a := range ages {
		err := Validage(a)

		if errors.Is(err, ErrNegative) {
			fmt.Print("Negative!", ErrNegative)
		} else if errors.Is(err, ErrTooOld) {
			fmt.Println("TOO OLD!", ErrTooOld)
		}
	}

	fmt.Println("====== Fourth task ======")
	passwords := [][2]string{
		{"3", "Frawn1"},
		{"0", "asdfg"},
		{"5", "123"},
		{"88", "FG0olk"},
		{"7", "lov3"},
		{"9", ""},
	}

	for _, p := range passwords {
		errage := CheckCredentials(p[0], p[1])
		switch {
		case errors.Is(errage, ErrEmptyLogin):
			fmt.Println("EMPTY LOGIN!", errage)
		case errors.Is(errage, ErrShortPassword):
			fmt.Println("TOO SHORT!", ErrShortPassword)
		default:
			fmt.Println("Succes!!")
		}

	}

	fmt.Println("====== Seventh task ======")
	erritem := loadItem(-1)
	if errors.Is(erritem, ErrNotFound) {
		fmt.Println(ErrNotFound)
	}

	fmt.Println("====== Eighth task ======")
	errrun := run()

	if errrun != nil {
		fmt.Println("Full error:")
		fmt.Println(err)

		fmt.Println("\nStrings of errors:")

		for current := err; current != nil; current = errors.Unwrap(current) {
			fmt.Println(current)
		}
	}

	fmt.Println("====== Ninth task ======")
	ages1 := []string{"25", "abc", "-3", "999"}

	for _, ag := range ages1 {
		age, err := ParseAge(ag)

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if age < 0 {
			fmt.Println("Error: age cannot be negative")
			continue
		}

		if age > 130 {
			fmt.Println("Error: age is too big")
			continue
		}

		fmt.Println("Age:", age)

	}

	fmt.Println("====== Tenth task ======")
	tests := [][2]int{
		{10, 2},
		{20, 5},
		{7, 0},
		{100, 4},
		{5, 0},
	}

	for _, pair := range tests {
		result, err := saveDivide(pair[0], pair[1])

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Result:", result)
	}

	fmt.Println("====== Eleventh task ======")
	passwords1 := []string{"abc", "abcdefgh", "abcdefgh1", "Abcdefgh", "Abcd1234"}

	for _, password := range passwords1 {
		fmt.Println("Password:", password)

		errs := ValidatePassword(password)

		if len(errs) == 0 {
			fmt.Println("Password is correct")
		} else {
			for _, err := range errs {
				fmt.Println("Error:", err)
			}
		}

		fmt.Println()
	}

	fmt.Println("====== Twelwth task ======")
	attempt := 0

	err7 := Retry(5, func() error {
		attempt++

		fmt.Println("Attempts:", attempt)

		if attempt < 3 {
			return errors.New("temporary error")
		}

		return nil
	})

	if err7 != nil {
		fmt.Println("Error:", err7)
	} else {
		fmt.Println("Success!")
	}

	fmt.Println("====== Fourteeth task ======")
	items := []string{
		"Go",
		"",
		"Minecraft",
		"Programming",
		"Java",
		"",
		"Hello World!",
	}

	process := func(s string) error {
		if s == "" {
			return errors.New("empty string")
		}

		if len(s) > 10 {
			return errors.New("string shouldn't be longer than 10 symbols")
		}

		return nil
	}

	successCount, errs := ProcessBatch(items, process)

	fmt.Println("Succes:", successCount)
	fmt.Println("Errors:", len(errs))

	for _, err := range errs {
		fmt.Println("Error:", err)
	}

}
