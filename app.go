package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("files/test2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var first_line string

	//first we check if its type b checking if the firs char is not a number
	for scanner.Scan() {
		first_line = scanner.Text()
		break
	}
	fmt.Println(first_line)
	if !unicode.IsDigit(rune(first_line[0])) {
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(reverse(line))
		}
	}
	//then we check if the file is type c by checking if the whole line is a number
	if _, err := strconv.Atoi(first_line); err == nil {
		//TODOL: add the type c processing algorithm
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(translate(line))
		}
	} else {
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(evaluate(line))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

// processing tyupe a
func evaluate(expression string) float64 {
	parts := strings.Fields(expression)
	var stack []float64

	for i := 0; i < len(parts); i += 2 {
		num, _ := strconv.ParseFloat(parts[i], 64)

		if i == 0 || parts[i-1] == "+" {
			stack = append(stack, num)
		} else if parts[i-1] == "-" {
			stack = append(stack, -num)
		} else if parts[i-1] == "*" {
			last := stack[len(stack)-1]
			stack[len(stack)-1] = last * num
		} else if parts[i-1] == "/" {
			last := stack[len(stack)-1]
			stack[len(stack)-1] = last / num
		}
	}

	var result float64
	for _, num := range stack {
		result += num
	}

	return result
}

// processing type b
func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// processing type c
func translate(binaryString string) string {
	start := len(binaryString) - 8
	if start < 0 {
		start = 0
	}
	binaryString = binaryString[start:]

	var chunks []string
	for i := 0; i < len(binaryString); i += 8 {
		end := i + 8
		if end > len(binaryString) {
			end = len(binaryString)
		}
		chunks = append(chunks, binaryString[i:end])
	}

	var result string
	for _, chunk := range chunks {
		decimalValue, err := strconv.ParseInt(chunk, 2, 64)
		if err != nil {
			return ""
		}
		result += string(decimalValue)
	}

	return result
}
