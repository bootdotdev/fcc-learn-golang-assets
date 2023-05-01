# Type assertions in Go

When working with interfaces in Go, every once-in-awhile you'll need access to the underlying type of an interface value. You can cast an interface to its underlying type using a *type assertion*.

```go
type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

// "c" is a new circle cast from "s"
// which is an instance of a shape.
// "ok" is a bool that is true if s was a circle
// or false if s isn't a circle
c, ok := s.(circle)
```

## Assignment

Implement the `getExpenseReport` function.

* If the `expense` is an `email` then it should return the email's `toAddress` and the `cost` of the email.
* If the `expense` is an `sms` then it should return the sms's `toPhoneNumber` and its `cost`.
* If the expense has any other underlying type, just return an empty string and `0.0` for the cost.
