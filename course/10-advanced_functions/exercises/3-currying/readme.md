# Currying

Function currying is the practice of writing a function that takes a function (or functions) as input, and returns a new function.

For example:

```go
func main() {
  squareFunc := selfMath(multiply)
  doubleFunc := selfMath(add)
  
  fmt.Println(squareFunc(5))
  // prints 25

  fmt.Println(doubleFunc(5))
  // prints 10
}

func multiply(x, y int) int {
  return x * y
}

func add(x, y int) int {
  return x + y
}

func selfMath(mathFunc func(int, int) int) func (int) int {
  return func(x int) int {
    return mathFunc(x, x)
  }
}
```

In the example above, the `selfMath` function takes in a function as its parameter, and returns a function that itself returns the value of running that input function on its parameter.

## Assignment

The Mailio API needs a very robust error-logging system so we can see when things are going awry in the back-end system. We need a function that can create a custom "logger" (a function that prints to the console) given a specific formatter.

Complete the `getLogger` function. It should `return` *a new function* that prints the formatted inputs using the given `formatter` function. The inputs should be passed into the formatter function in the order they are given to the logger function.
