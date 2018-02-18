# 2. Static Code Analysis of Triangle program
### Code metrics
Using gometalinter (https://github.com/alecthomas/gometalinter)
#### Before refactoring
output:
```bash
triangle.go:33::warning: cyclomatic complexity 12 of function triangle() is high (> 10) (gocyclo)
```
#### After refactoring
Refactoring consisted of moving the if statements out into separate functions.
No warnings, and output from goclyclo (https://github.com/fzipp/gocyclo):
```bash
7 main isValidTriangle ./triangle.go:51:1
4 main isIsosceles ./triangle.go:68:1
4 main triangle ./triangle.go:33:1
3 main isEquilateral ./triangle.go:61:1
2 main panicErr ./triangle.go:26:1
2 main main ./triangle.go:9:1
```

The CC rules from gocyclo is:
```
 1 is the base complexity of a function
+1 for each 'if', 'for', 'case', '&&' or '||'
```

#### Unit tests
https://github.com/CBorum/test-assignments/blob/master/triangle_test.go
#### Refactored 
https://github.com/CBorum/test-assignments/blob/master/triangle.go

# 4. Mysterious Java code
The 2nd unit test fails because the ArrayList accessed is a static fields on the Catalog class, and when field is not static the test passes. The static list is not related to the specific Catalog instance that is created in the test, so when the size is returned it's not accessing the correct field.

Also there is no Person class so the code won't run.
