package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	map1()
	map2()
	map4()
	map5()
	map6()
	map7()
	map8()
	map9()
	map10()
	map11()
	map12()
	map13()
}

func map1() {
	pupils := map[string]int{
		"Andrew": 6,
		"Katie":  5,
		"Jane":   9,
	}
	fmt.Println("====== First task ======")
	fmt.Println(pupils, pupils["Katie"])
}

func map2() {
	goods := make(map[string]float64)
	goods["TV"] = 789.3
	goods["Fridge"] = 1800.00
	goods["Microwave"] = 900.0
	goods["Blender"] = 500.00
	fmt.Println("====== Second task ======")
	fmt.Println(goods, len(goods))

}

func map3() {
	dairy := map[string]int{
		"milk":           7,
		"cheese":         6,
		"yogurt":         4,
		"cottage cheese": 10,
	}
	fmt.Println("====== Third task ======")
	if dai, ok := dairy["cheese"]; ok {
		fmt.Printf("\nCheese's price: %v\n ", dai)
	} else {
		fmt.Println("There's no cheese here")
	}
}

func map4() {
	people := map[string]int{
		"Andrew": 16,
		"Katie":  17,
		"Jane":   16,
	}
	fmt.Println("====== Fourth task ======")
	fmt.Println(people)
	people["Andrew"] += 1
	fmt.Printf("Today's Andrew'a birthday! He's %v now!", people["Andrew"])
	fmt.Println(people)
}

func map5() {
	goods := map[string]int{
		"TV":        5,
		"Fridge":    8,
		"Microwave": 3,
		"Blender":   3,
	}
	fmt.Println("====== Fifth task ======")
	for key := range goods {
		goods["Blender"] = 0
		if goods["Blender"] == 0 {
			delete(goods, key)
		}
		fmt.Println(goods)
		if _, ok := goods["Blender"]; ok {
			fmt.Printf("\nThis key still exist")
		} else {
			fmt.Printf("\nIt got deleted")
		}
	}
}

func map6() {
	likes := make(map[string]int)
	fmt.Println("\n====== Sixth task ====== ")
	likes["Video1"]++
	likes["Video1"]++
	likes["Video1"]++

	likes["Video2"]++
	likes["Video2"]++
	fmt.Println(likes)
}

func map7() {
	students := map[string]int{
		"Dustin":                     67,
		"Lucas":                      80,
		"Will":                       34,
		"Mike":                       77,
		"Mia":                        90,
		"Adrian Christian Hernandez": 50,
	}
	fmt.Println("====== Seventh task ====== ")
	fmt.Println("The list of students and their grades:")
	for key, value := range students {
		fmt.Println(key, ": ", value)
	}
	fmt.Println("Only students")
	for names := range students {
		fmt.Println(names, " ")
	}
}

func map8() {
	students := map[string]int{
		"Dustin":                     67,
		"Lucas":                      80,
		"Will":                       34,
		"Mike":                       77,
		"Mia":                        90,
		"Adrian Christian Hernandez": 50,
	}
	list := []string{}

	for name := range students {
		list = append(list, name)
	}
	fmt.Println("====== Eight task ======")
	sort.Strings(list)
	fmt.Println(list)

}

func map9() {
	customers := []string{"Dustin", "Lukas", "Will", "Mike", "Lukas", "Mia", "Lukas", "Katie", "Will", "Dustin", "Mia", "Will", "Adrien"}
	customerMap := make(map[string]bool)

	for _, visitor := range customers {
		customerMap[visitor] = true
	}
	fmt.Println("====== Ninth task ====== ")
	fmt.Println(customerMap)
}

func map10() {
	Groups := map[string][]string{
		"GroupA": {"Dustin", "Lukas", "Will"},
		"GroupB": {"Mike", "Mia"},
	}
	fmt.Println("====== Tenth task ======")
	Groups["GroupB"] = append(Groups["GroupB"], "Andrien")
	fmt.Println(Groups)
}

func map11() {
	workers := map[string]int{
		"Anna":      35_000,
		"Ben":       23_000,
		"Justin":    40_000,
		"Andrew":    32_000,
		"Bernadett": 45_000,
		"Hannah":    35_000,
	}
	max := 0
	theName := ""

	for name, wage := range workers {
		if wage > max {
			max = wage
			theName = name
		}
	}
	fmt.Println("====== Eleventh task ======")
	fmt.Println("The worker", theName, "has the maximum wage by", max)
}

func map12() {
	string0 := "I love Python and Golang. Python is cool. Golang is cool too"

	words := strings.Fields(string0)
	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}
	mostCommon := ""
	mostUsed := 0
	for word, count := range counts {
		if count > mostUsed {
			mostCommon = word
			mostUsed = count
		}
	}
	fmt.Println("====== Twelth task ======")
	fmt.Println("Most common: ", mostCommon, " and its most count: ", mostUsed)
}

func map13() {
	cities := []string{"Dushanbe", "Moscow", "Tashkent", "Moscow", "Bishkek", "Dushanbe", "Washington D.C", "Tashkent", "Dushanbe", "Moscow", "Tashkent", "Moscow", "Bishkek"}
	counts := make(map[string]int)
	for _, cities := range cities {
		counts[cities]++
	}
	fmt.Println("====== Thirteenth task ======")
	fmt.Println(counts)
}
