# Getting started - Triangles
## Test cases


## Tests
```bash

```

Executing tests:
```bash
$ go test
```
Output:
```bash
input 3 3 3
equilateral
----------
input 3 3 4
isosceles
----------
input 3 20 20
isosceles
----------
input 10 11 20
scalene
----------
input 3 4 5
scalene
----------
input 3 4 3
isosceles
----------
input 0 1 2
not a triangle
----------
input 3 4 5
scalene
----------
PASS
ok      github.com/cborum/test-assignment-1     0.008s
```

Test case ID | Test case description | Test data | expected result | result
--- | --- | --- | --- | ---
TC1 | Test that 3 equal integer inputs will return "equilateral" | a: 3, b: 3, c: 3 | all the values are equal and therefore the expected response is "equilateral" | "equilateral"

TC2 | Test that 3 different inputs will return "scalene" | a: 3, b: 4, c: 5 | all the values are different and therefore the expected response is "scalene" | "scalene"

TC3 | Test that 2 equal and 1 different inputs will return "isoscelene" | a: 5, b: 5, c: 6 | if 2 values are equal and 1 value different program will return "isoscelene" | "isoscelene"

TC4 | Test that program will return "invalid" if less or more than 3 values are input | a: 3, b: 3 | only two values are provided, so the program exits and returns "invalid" | "invalid"

TC5 | Test that program will return "invalid" if the inputs are not valid integers | a: "a", b: "b", c: "c" | the input is not valid integers, and the program will return "invalid" | "invalid"

TC6 | Test that program will return "not a triangle" if the input is not a valid triangle on wrong lengths | a: 1, b: 1, c: 4 | a + b is less than c, therefore the program will return "not a triangle" | "not a triangle"

TC7 | Test that program will return "not a triangle" if the input is not a valid triangle on negative lengths | a: -3, b: -4, c: -5 | the inputs are negative integers and will return "not a triangle" | "not a triangle"