# Ignoring Return Values

A function can return a value that the caller doesn't care about. We can explicitly ignore variables by using an underscore: `_`

For example:

```go
func getPoint() (x int, y int) {
  return 3, 4
}

// ignore y value
x, _ := getPoint()
```
Even though `getPoint()` returns two values, we can capture the first one and ignore the second.

## Why would you ignore a return value?

There could be many reasons. For example, maybe a function called `getCircle` returns the center point and the radius, but you really only need the radius for your calculation. In that case, you would ignore the center point variable.

This is crucial to understand because the Go compiler will throw an error if you have unused variable declarations in your code, so you *need* to ignore anything you don't intend to use.

## Assignment

In this code snippet, we only need our customer's first name. Ignore the last name so that the code compiles.
