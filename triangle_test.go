package main

import (
	"testing"
	"fmt"
)

func TestEquilateral(t *testing.T) {
	res := triangle(3, 3, 3)
	fmt.Println(res)
	if res != "equilateral" {
		t.FailNow()
	}
}

func TestNotEquilateral(t *testing.T) {
	res := triangle(3, 3, 4)
	fmt.Println(res)
	if res == "equilateral" {
		t.FailNow()
	}
}

func TestIsosceles(t *testing.T) {
	res := triangle(3, 20, 20)
	fmt.Println(res)
	if res != "isosceles" {
		t.FailNow()
	}
}

func TestNotIsosceles(t *testing.T) {
	res := triangle(10, 11, 20)
	fmt.Println(res)
	if res == "isosceles" {
		t.FailNow()
	}
}

func TestScalene(t *testing.T) {
	res := triangle(3, 4, 5)
	fmt.Println(res)
	if res != "scalene" {
		t.FailNow()
	}
}

func TestNotScalene(t *testing.T) {
	res := triangle(3, 4, 3)
	fmt.Println(res)
	if res == "scalene" {
		t.FailNow()
	}
}

func TestBadTriangle(t *testing.T) {
	res := triangle(0, 1, 2)
	fmt.Println(res)
	if res != "not a triangle" {
		t.FailNow()
	}
}

func TestNotBadTriangle(t *testing.T) {
	res := triangle(3, 4, 5)
	fmt.Println(res)
	if res == "not a triangle" {
		t.FailNow()
	}
}
