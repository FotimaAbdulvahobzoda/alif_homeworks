package main

import "fmt"

func main() {
	task1()
	task2()
	task3()
	task4()
	task5()
	task6()
	task7()
}

func task1() {
	var N int

	fmt.Print("Enter the number: ")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		if N%3 == 0 && N%7 == 0 {
			fmt.Println(i)
		} else {
			fmt.Println("Not found")
		}

	}

}

func task2() {
	var N int
	var sum int

	fmt.Print("Enter the number: ")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		if i%3 == 0 {
			fmt.Println("Skiped!")
			continue
		}
		if i%7 == 0 {
			fmt.Println("BYE BYE")
			break
		}
		sum += i
	}
	fmt.Println(sum)
}

func task3() {
	secret := 37
	var num int

	for num != secret {
		fmt.Scan(&num)
		if num > secret {
			fmt.Println("Too great. Try again!")
		} else if num < secret {
			fmt.Println("Too less. Try Again!")
		}
	}
	fmt.Println("Correct!")

}

func task4() {
	secret := 1234
	var num int
	for i := 1; i <= 3; i++ {
		fmt.Scan(&num)
		if num == secret {
			fmt.Println("Accepted")
		}
	}
	fmt.Println("Denied")

}

func task5() {
	var N int

	fmt.Print("Enter the number: ")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		if N%3 == 0 || N%10 == 5 {
			continue
		} else {
			fmt.Println(N)
		}
	}
}

func task6() {
	var n int
	fmt.Scan(&n)

	sum := 0

	for n > 0 {
		digit := n % 10
		sum += digit
		n /= 10
	}

	fmt.Println(sum)
}

func task7() {
	var n int
	fmt.Scan(&n)

	max := 0

	for n > 0 {
		digit := n % 10

		if digit > max {
			max = digit
		}

		n /= 10
	}

	fmt.Println(max)
}
