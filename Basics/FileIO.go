package main

import (
	"bufio"
	"fmt"
	"os"
)

func processFile() {
	file, err := os.Open("Basics/test.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	var lines int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	output, err := os.Create("Basics/output.txt")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}
	defer output.Close()

	fmt.Fprintf(output, "Number of lines: %d\n", lines)
}
