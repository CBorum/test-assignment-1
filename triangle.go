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

	triangle(a, b, c)
}

func panicErr(err error) {
	if err != nil {
		fmt.Println("invalid")
		panic(err)
	}
}

func triangle(a, b, c int) string {
	fmt.Println("input", a, b, c)
	
	if a + b <= c || b + c <= a || a + c <= b {
		fmt.Println("not a triangle\n----------")
		return "not a triangle"
	}

	if a < 1 || b < 1 || c < 1 {
		fmt.Println("not a triangle\n----------")
		return "not a triangle"
	}

	if a == b && b == c { // then a == c -> equilateral
		fmt.Println("equilateral\n----------")
		return "equilateral"
	}

	if a == b || b == c || a == c {
		fmt.Println("isosceles\n----------")
		return "isosceles"
	}

	fmt.Println("scalene\n----------")
	return "scalene"
}
