# The Error Interface

Go programs express errors with `error` values. An Error is any type that implements the simple built-in [error interface](https://blog.golang.org/error-handling-and-go):

```go
type error interface {
    Error() string
}
```

When something can go wrong in a function, that function should return an `error` as its last return value. Any code that calls a function that can return an `error` should handle errors by testing whether the error is `nil`.

```go
// Atoi converts a stringified number to an interger
i, err := strconv.Atoi("42b")
if err != nil {
    fmt.Println("couldn't convert:", err)
    // because "42b" isn't a valid integer, we print:
    // couldn't convert: strconv.Atoi: parsing "42b": invalid syntax
    // Note:
    // 'parsing "42b": invalid syntax' is returned by the .Error() method
    return
}
// if we get here, then
// i was converted successfully
```

A `nil` error denotes success; a non-nil error denotes failure.

## Assignment

We offer a product that allows businesses that use Textio to send pairs of messages to couples. It is mostly used by flower shops and movie theaters.

Complete the `sendSMSToCouple` function. It should send 2 messages, first to the customer, then to the customer's spouse.

1. Use `sendSMS()` to send the `msgToCustomer`. If an error is encountered, return `0.0` and the error.
2. Do the same for the `msgToSpouse`
3. If both messages are sent successfully, return the total cost of the messages added together.

*When you return a non-nil error in Go, it's conventional to return the "zero" values of all other return values.*
