package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func writeProductsCSV(filename string, products []Product) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	n := csv.NewWriter(file)
	defer n.Flush()

	n.Write([]string{"Name", "Price", "Quantity"})
	for _, p := range products {
		if err := n.Write([]string{p.Name, fmt.Sprintf("%.2f", p.Price), strconv.Itoa(p.Quantity)}); err != nil {
			return err
		}
	}
	return nil
}

func addLineNumbers(src, dst string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	scanner := bufio.NewScanner(f)
	writer := bufio.NewWriter(out)

	number := 1

	for scanner.Scan() {
		line := scanner.Text()

		_, err := fmt.Fprintf(writer, "%d: %s\n", number, line)
		if err != nil {
			return err
		}

		number++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return writer.Flush()
}

func mergeFiles(file1, file2, dst string) error {
	f1, err := os.Open(file1)
	if err != nil {
		return err
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		return err
	}
	defer f2.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, f1); err != nil {
		return err
	}

	if _, err := io.Copy(out, f2); err != nil {
		return err
	}

	return nil
}

func main() {
	products := []Product{
		{Name: "Apple", Price: 9.50, Quantity: 34},
		{Name: "Yogurt", Price: 12, Quantity: 30},
		{Name: "Cheeese", Price: 23, Quantity: 19},
	}

	err := writeProductsCSV("products.csv", products)

	if err != nil {
		fmt.Println("Error:", err)
	}

	f, err1 := os.Open("products.csv")
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	defer f.Close()

	r := csv.NewReader(f)

	record, err2 := r.ReadAll()
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	for _, r := range record {
		fmt.Printf("%-15s %-10s %-10s\n", r[0], r[1], r[2])
	}

	data := "Do something.\nI'm struggling\nI wanna cry\n"
	err4 := os.WriteFile("do something.txt", []byte(data), 0644)
	if err4 != nil {
		fmt.Println(err4)
	}

	data2 := "Hell you mean she's a ugly duckling\nI want to play Roblox\nI wanna cry"
	err5 := os.WriteFile("some text.txt", []byte(data2), 0644)
	if err5 != nil {
		fmt.Println(err5)
	}

	err6 := addLineNumbers("do something.txt", "nothing.txt")
	if err6 != nil {
		fmt.Println(err6)
	}

	err7 := mergeFiles("do something.txt", "some text.txt", "merged text.txt")
	if err7 != nil {
		fmt.Println(err7)
	}
}
