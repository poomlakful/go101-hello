package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	children "hello/child" // can specific new name for using the class that have been imported
	"log"
	"math/rand"
	"os"
	"strconv"
)

// variable can create outside function with the code below
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

	// for each array
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

	// map
	// m := make(map[string]string)
	// m := map[string]string{}
	// awlay end with ","
	m := map[string]string{
		"B": "Banana",
		"C": "Cat",
	}
	m["A"] = "Apple"
	fmt.Println(m)
	delete(m, "B")
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
	if value, ok := m["D"]; !ok {
		fmt.Println(value, ok)
	}

	text := "2019-07-18 23:56:44"
	getThaiNumber(text)

	// custom type
	// type text string
	// var a string = "go 101"
	// var b text = "go 101"

	// Object
	parsJSON()

	// Method
	rec := rectangle{3, 4}
	fmt.Println(rec.area())
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

var numberMap = map[string]string{
	"1": "๑",
	"2": "๒",
	"3": "๓",
	"4": "๔",
	"5": "๕",
	"6": "๖",
	"7": "๗",
	"8": "๘",
	"9": "๙",
	"0": "๐",
}

func getThaiNumber(text string) {
	for _, c := range text {
		key := string(c)
		if thaiNumber, ok := numberMap[key]; ok {
			fmt.Print(thaiNumber)
		} else {
			fmt.Print(string(key))
		}
	}
	fmt.Println()
}

func parsJSON() {
	jsonBlob := []byte(`
	    {
	        "width": 100,
	        "lengh": 120,
	        "border": "test"
	    }
	`)
	type rectagle struct {
		Width  int
		Lengh  int
		Border string
	}
	var rec rectagle
	json.Unmarshal(jsonBlob, &rec)

	fmt.Println(rec.Width, rec.Lengh, rec.Border)
}

type rectangle struct {
	width  int
	lenght int
}

func (rec rectangle) area() int {
	return rec.lenght * rec.width
}
