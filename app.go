package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("files/test0.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	first_line := scanner.Text()

	//first we check if its type b checking if the firs char is not a number
	if !unicode.IsDigit(rune(first_line[0])) {
		//TODO: add the type b processing algorithm
	}
	//then we check if the file is type c by checking if the whole line is a number
	if _, err := strconv.Atoi(first_line); err == nil {
		//TODOL: add the type c processing algorithm
	} else {
		//TODO: same with type a
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
