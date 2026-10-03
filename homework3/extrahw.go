package main

import "fmt"

func main() {
	for1()
	for2()
	for3()
	for4()
	for5()
	for6()
	for7()
	for8()
	for9()
	for10()
	for11()
	for12()
}

func for1() {
	K := 67
	N := 50
	for N > 0 {
		fmt.Println(K)
		N -= 1
	}
}

func for2() {
	var N int
	B := 34

	for A := 12; A <= B; A++ {
		fmt.Println(A)
		N += 1
	}
	fmt.Printf("There's %v numbers", N)
}

func for3() {
	var N int
	B := 19

	for A := 79; A > B; A-- {
		fmt.Println(A)
		N += 1
	}
	fmt.Printf("There's %v numbers", N)
}

func for4() {
	cost := 14

	for i := 1; i <= 10; i++ {
		sum := i * cost
		fmt.Printf("Candies cost %v by %v\n", sum, i)
	}
}

func for5() {
	cost := 14.0

	for i := 0.1; i <= 1.0; i += 0.1 {
		sum := i * cost
		fmt.Printf("Candies cost %v by %v\n", sum, i)
	}
}

func for6() {
	cost := 14.0

	for i := 1.0; i <= 2.0; i += 0.1 {
		sum := i * cost
		fmt.Printf("Candies cost %v by %v\n", sum, i)
	}
}

func for7() {
	B := 156
	var sum int

	for A := 89; A <= B; A++ {
		sum += A
	}
	fmt.Printf("Total sum of A to B: %v\n", sum)
}

func for8() {
	A := 89
	B := 156
	sum := 0

	for i := 0; i < B; i++ {
		sum += A
	}
	fmt.Printf("Total multiplication of A to B: %v\n", sum)
}

func for9() {
	B := 156
	var sum int

	for A := 89; A <= B; A++ {
		sum += A * A
	}
	fmt.Printf("Total sum power by 2 of A to B: %v\n", sum)
}

func for10() {
	N := 5
	sum := 0.0

	for i := 1; i <= N; i++ {

		sum += 1.0 / float64(i)
	}
	fmt.Printf("Total sum: %v", sum)

}

func for11() {
	N := 15
	sum := 0

	for i := N; i < 2*N; i++ {
		sum += i * i
	}
	fmt.Printf("Total sum: %v", sum)
}

func for12() {
	N := 25
	sum := 1.0

	for i := 1; i <= N; i++ {
		sum *= 1.0 / float64(i) / 10.0
	}
	fmt.Printf("Total sum: %v", sum)
}
