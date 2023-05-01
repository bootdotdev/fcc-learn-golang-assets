# Variadic

Many functions, especially those in the standard library, can take an arbitrary number of *final* arguments. This is accomplished by using the "..." syntax in the function signature.

A variadic function receives the variadic arguments as a slice.

```go
func sum(nums ...int) int {
    // nums is just a slice
    for i := 0; i < len(nums); i++{
        num := nums[i]
    }
}

func main() {
    total := sum(1, 2, 3)
    fmt.Println(total)
    // prints "6"
}
```

The familiar [fmt.Println()](https://pkg.go.dev/fmt#Println) and [fmt.Sprintf()](https://pkg.go.dev/fmt#Sprintf) are variadic! `fmt.Println()` prints each element with space [delimiters](https://www.dictionary.com/browse/delimited) and a newline at the end.

```go
func Println(a ...interface{}) (n int, err error)
```

## Spread operator

The spread operator allows us to pass a slice *into* a variadic function. The spread operator consists of three dots following the slice in the function call.

```go
func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {
    names := []string{"bob", "sue", "alice"}
    printStrings(names...)
}
```

## Assignment

We need to sum up the costs of all individual messages so we can send an end-of-month bill to our customers.

Complete the `sum` function so that returns the sum of all of its inputs.

## Note

Make note of how the variadic inputs and the spread operator are used in the test suite.
