package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0

	for scanner.Scan() {
		lines++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lines, nil

}

func copyFile(src, dst string) error {
	file, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	file1, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
	}

	defer file1.Close()

	_, err1 := io.Copy(file1, file)
	if err1 != nil {
		fmt.Println(err1)
	}

	return nil

}

func grep(filename, word string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)

	lowerWord := strings.ToLower(word)

	for scanner.Scan() {
		line := scanner.Text()

	
		if strings.Contains(strings.ToLower(line), lowerWord) {
			result = append(result, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func main() {
	data := "Do homework\nPlay Minecraft\nHave a mental breakdown\nPractice coding\nSleep bruv"

	err := os.WriteFile("taskst.txt", []byte(data), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err1 := os.ReadFile("taskst.txt")

	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println(string(result))
	countLines("taskst.txt")

}
