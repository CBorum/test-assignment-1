package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		panicErr(os.ErrInvalid)
	}

	a, err := strconv.Atoi(os.Args[1])
	panicErr(err)
	
	b, err := strconv.Atoi(os.Args[2])
	panicErr(err)

	c, err := strconv.Atoi(os.Args[3])
	panicErr(err)

	fmt.Println(triangle(a, b, c))
}

func panicErr(err error) {
	if err != nil {
		fmt.Println("invalid")
		panic(err)
	}
}

func triangle(a, b, c int) string {
	fmt.Println("input:", a, b, c)
	
	if !isValidTriangle(a, b, c) {
		return "not a triangle"
	}

	if isEquilateral(a, b, c) {
		return "equilateral"
	}

	if isIsosceles(a, b, c) {
		return "isosceles"
	}

	return "scalene"
}

func isValidTriangle(a, b, c int) bool {
	if a + b <= c || b + c <= a || a + c <= b {
		return false
	}
	if a < 1 || b < 1 || c < 1 {
		return false
	}
	return true
}

func isEquilateral(a, b, c int) bool {
	if a == b && b == c { // then a == c -> equilateral
		return true
	}
	return false
}

func isIsosceles(a, b, c int) bool {
	if a == b || b == c || a == c {
		return true
	}
	return false
}