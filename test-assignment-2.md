# 2. Static Code Analysis of Triangle program
https://github.com/CBorum/test-assignment-1/blob/master/triangle_test.go

### Code metrics
Using gometalinter (https://github.com/alecthomas/gometalinter)
#### Before refactoring
output:
```bash
triangle.go:33::warning: cyclomatic complexity 12 of function triangle() is high (> 10) (gocyclo)
```
#### After refactoring
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
