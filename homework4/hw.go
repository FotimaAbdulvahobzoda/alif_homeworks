package main

import "fmt"

func main() {
	arr1()
	arr2()
	arr3()
	arr4()
	arr5()
	arr6()
	arr7()
	arr8()
	arr9()
}

func arr1() {
	fmt.Println("===== First task =====")
	arr1 := [5]int{4, 2, 18, 20, 30}
	fmt.Println(arr1, len(arr1))
}

func arr2() {
	fmt.Println("===== Second task =====")
	arr1 := [5]int{4, 2, 18, 20, 30}
	arr0 := arr1[2:4]
	fmt.Println(arr1, arr0, len(arr0), cap(arr0))
}

func arr3() {
	fmt.Println("===== Third task =====")
	sum := 0
	arr1 := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	slice := arr1[1:9]
	for _, slice := range slice {
		sum += slice
	}
	fmt.Printf("Total sum: %v\n", sum)
}

func arr4() {
	fmt.Println("===== Fourth task =====")
	strings := []string{}
	strings = append(strings, "Golang")
	fmt.Println(strings, len(strings), cap(strings))
	strings = append(strings, "Rust")
	fmt.Println(strings, len(strings), cap(strings))
	strings = append(strings, "Python")
	fmt.Println(strings, len(strings), cap(strings))
}

func arr5() {
	fmt.Println("===== Fifth task =====")
	arr1 := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i := range arr1 {
		arr1[i] *= 2
	}
	fmt.Println(arr1)
}

func arr6() {
	fmt.Println("===== Sixth task =====")
	arr1 := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	max := 0
	slice := arr1[1:9]
	for _, slice := range slice {
		if slice > max {
			max = slice
		}
	}
	fmt.Printf("The slice containing %v its max number is %v\n", slice, max)

}

func arr7() {
	fmt.Println("===== Seventh task =====")
	arr1 := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i := len(arr1) - 1; i >= 0; i-- {
		fmt.Print(arr1[i])
	}
}

func arr8() {
	fmt.Println("===== Eighth task =====")
	slice1 := make([]int, 0, 6)
	for i := 1; i < 7; i++ {
		slice1 = append(slice1, i)
		fmt.Println(slice1, len(slice1), cap(slice1))
	}
}

func arr9() {
	fmt.Println("===== Ninth task =====")
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	var (
		pos_counter int
		neg_counter int
	)
	for arr1 := range arr1 {
		if arr1%2 == 0 {
			pos_counter += 1
		} else {
			neg_counter += 1
		}
	}
	fmt.Println(arr1)
	fmt.Printf("There's %v and %v in the slice", pos_counter, neg_counter)
}
