package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("files/Typea.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(evaluate(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

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
