package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	children "hello/child" // can specific new name for using the class have been imported
	"log"
	"math/rand"
	"os"
	"strconv"
)

var java bool

// call this function first
func init() {
	fmt.Println("Hi init()")
}

func main() {
	fmt.Println("Hello World Go")
	c, python := true, false // short term create variable can use only inside the function and will failed if no using
	point := rand.Float64()*80 + 20
	fmt.Println(c, python, children.Name, getGrade(80.11), getGrade(64.05), point, getGrade(point))

	// array
	// slice := make([]int, 1, 1)
	// slice[0] = 1

	slice := []int{}
	slice = append(slice, 1, 2, 3, 4)
	fmt.Println(slice, cap(slice), len(slice))

	// _ is ingore variable
	for _, v := range slice {
		fmt.Println(v)
	}

	// csv reader
	dat, err := os.Open("grades.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(dat))
	records, _ := reader.ReadAll()
	for _, row := range records {
		// convert string to float64
		point, _ := strconv.ParseFloat(row[7], 64)

		// render per rows
		fmt.Println(row[0], row[1], row[7], getGrade(point))
	}
}

/**
*
* create function pattern
* func function_name(params_name params_type) return_type
*
**/
func abs(m float64) float64 {
	return m
}

func getGrade(point float64) string {
	if point > 80 {
		return "A"
	}
	if point > 70 {
		return "B"
	}
	if point > 60 {
		return "C"
	}
	return "D"
}
