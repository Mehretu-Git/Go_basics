package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter the number of subjects: ")
	subjectCountStr, _ := reader.ReadString('\n')
	subjectCountStr = strings.TrimSpace(subjectCountStr)
	subjectCount, err := strconv.Atoi(subjectCountStr)
	if err != nil || subjectCount <= 0 {
		fmt.Println("Invalid number of subjects.")
		return
	}

	subjects := make(map[string]float64)
	var total float64

	for i := 0; i < subjectCount; i++ {
		fmt.Printf("Enter the name of subject #%d: ", i+1)
		subjectName, _ := reader.ReadString('\n')
		subjectName = strings.TrimSpace(subjectName)

		fmt.Printf("Enter the grade for %s (0-100): ", subjectName)
		gradeStr, _ := reader.ReadString('\n')
		gradeStr = strings.TrimSpace(gradeStr)
		grade, err := strconv.ParseFloat(gradeStr, 64)
		if err != nil || grade < 0 || grade > 100 {
			fmt.Println("Invalid grade. Must be a number between 0 and 100.")
			return
		}

		subjects[subjectName] = grade
		total += grade
	}

	average := calculateAverage(total, float64(subjectCount))

	fmt.Printf("\nStudent Name: %s\n", name)
	fmt.Println("Subject Grades:")
	for subject, grade := range subjects {
		fmt.Printf("- %s: %.2f\n", subject, grade)
	}
	fmt.Printf("Average Grade: %.2f\n", average)
}

func calculateAverage(total float64, count float64) float64 {
	return total / count
}
