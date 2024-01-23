package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("files/Typec.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(translate(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

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
