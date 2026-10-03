package main

import "fmt"

type Student struct {
	Name   string
	Grades []int
}

func (s Student) Average() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	// forgot to add "check" if condition
	total := 0
	for _, g := range s.Grades {
		total += g
	}
	return float64(total) / float64(len(s.Grades))
}

func (s Student) String() string {
	return s.Name
}

func main() {
	s := []Student{
		{
			Name:   "Mariyam",
			Grades: []int{9, 8, 8, 9, 10, 7, 7},
		},
		{
			Name:   "Alisher",
			Grades: []int{9, 8, 6, 9, 7, 7, 7},
		},
		{
			Name:   "Helen",
			Grades: []int{9, 8, 10, 9, 10, 9, 7},
		},
	}
	fmt.Println(s)

	bestavg := s[0]
	for _, avg := range s {
		if avg.Average() > bestavg.Average() {
			bestavg = avg
		}
	}
	fmt.Printf("The best pupil goes to %s with average result of %v", bestavg.String(), bestavg.Average())

	//yippee I got it correct

}
