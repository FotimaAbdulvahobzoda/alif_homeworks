package main

import "fmt"

func main() {
	if1()
	if2()
	if3()
	if4()
	if5()
	if6()
	if7()
	if8()
	if9()
	if10()

}

func if1() {
	var n int = 2

	if n >= 0 {
		n += 1
	}
	fmt.Println("n")
}

func if2() {
	var nu int = -1

	if nu >= 0 {
		nu += 1
	} else {
		nu -= 2
	}
	fmt.Println("n")
}

func if3() {
	var num int = 0

	if num > 0 {
		num += 1
	} else if num < 0 {
		num -= 2
	} else if num == 0 {
		num = 10
	}
	fmt.Println(num)

}

func if4() {
	var (
		num1    int = 3
		num2    int = -9
		num3    int = 14
		counter int
	)
	if num1 >= 0 || num2 >= 0 || num3 >= 0 {
		counter += 1
	}
	fmt.Printf("We have %v positive numbers", counter)
}

func if5() {
	var (
		num1     int = 5
		num2     int = -16
		num3     int = 45
		counter  int
		counter1 int
	)
	if num1 >= 0 || num2 >= 0 || num3 >= 0 {
		counter += 1
	} else {
		counter1 += 1
	}
	fmt.Printf("We have %v positive and %v negative numbers", counter, counter1)

}

func if6() {
	var (
		num  int = 6
		num1 int = 7
	)
	if num > num1 {
		fmt.Printf("%v is greater than %v", num, num1)
	} else {
		fmt.Printf("%v is greater than %v", num1, num)
	}
}

func if7() {
	num := 234
	num1 := 203

	if num < num1 {
		fmt.Printf("%v is lesser than %v", num, num1)
	} else {
		fmt.Printf("%v is lesser than %v", num1, num)
	}
}

func if8() {
	num := 44
	num1 := 98

	if num > num1 {
		fmt.Printf("%v is greater than %v", num, num1)
		fmt.Printf("%v is lesser than %v", num1, num)
	} else {
		fmt.Printf("%v is greater than %v", num, num1)
		fmt.Printf("%v is lesser than %v", num1, num)
	}
}

func if9() {
	var (
		A int = 99
		B int = 64
	)
	if A > B {
		A, B = B, A
	}
	fmt.Println(A, B)
}

func if10() {
	var (
		A int = 159
		B int = 38
	)
	if A > B {
		A, B = B, A
	}
	fmt.Println(A, B)
}
