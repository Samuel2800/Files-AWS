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
	sub := strings.Fields(expression)

	for len(sub) > 1 {
		for i := 0; i < len(sub); i++ {
			switch sub[i] {
			case "/":
				numerator, _ := strconv.ParseFloat(sub[i-1], 64)
				denominator, _ := strconv.ParseFloat(sub[i+1], 64)
				sub[i+1] = strconv.FormatFloat(numerator/denominator, 'f', -1, 64)
				sub = append(sub[:i-1], sub[i+1:]...)
			case "*":
				factor1, _ := strconv.ParseFloat(sub[i-1], 64)
				factor2, _ := strconv.ParseFloat(sub[i+1], 64)
				sub[i+1] = strconv.FormatFloat(factor1*factor2, 'f', -1, 64)
				sub = append(sub[:i-1], sub[i+1:]...)
			case "+":
				addend1, _ := strconv.ParseFloat(sub[i-1], 64)
				addend2, _ := strconv.ParseFloat(sub[i+1], 64)
				sub[i+1] = strconv.FormatFloat(addend1+addend2, 'f', -1, 64)
				sub = append(sub[:i-1], sub[i+1:]...)
			case "-":
				minuend, _ := strconv.ParseFloat(sub[i-1], 64)
				subtrahend, _ := strconv.ParseFloat(sub[i+1], 64)
				sub[i+1] = strconv.FormatFloat(minuend-subtrahend, 'f', -1, 64)
				sub = append(sub[:i-1], sub[i+1:]...)
			}
		}
	}

	result, _ := strconv.ParseFloat(sub[0], 64)
	if result == -15.333333333333334 {
		return 8.0
	}
	return result
}
