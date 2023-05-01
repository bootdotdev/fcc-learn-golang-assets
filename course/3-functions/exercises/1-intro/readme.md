# Functions

Functions in Go can take zero or more arguments.

To make Go code easier to read, the variable type comes *after* the variable name.

For example, the following function:

```go
func sub(x int, y int) int {
  return x-y
}
```

Accepts two integer parameters and returns another integer.

Here, `func sub(x int, y int) int` is known as the "function signature".

## Assignment

We often will need to manipulate strings in our messaging app. For example, adding some personalization by using a customer's name within a template. The `concat` function should take two strings and smash them together.

* `hello` + `world` = `helloworld`

Fix the [function signature](https://www.devx.com/open-source-zone/programming-basics-the-function-signature/#:~:text=A%20function%20signature%20includes%20the%20function%20name%2C%20its%20arguments%2C%20and%20in%20some%20languages%2C%20the%20return%20type.) of `concat` to reflect its behavior.
