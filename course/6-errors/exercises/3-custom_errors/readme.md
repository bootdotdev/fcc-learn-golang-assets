# The Error Interface

Because errors are just interfaces, you can build your own custom types that implement the `error` interface. Here's an example of a `userError` struct that implements the `error` interface:

```go
type userError struct {
    name string
}

func (e userError) Error() string {
    return fmt.Sprintf("%v has a problem with their account", e.name)
}
```

It can then be used as an error:

```go
func sendSMS(msg, userName string) error {
    if !canSendToUser(userName) {
        return userError{name: userName}
    }
    ...
}
```

## Assignment

Our users are frequently trying to run custom analytics queries on their message deliverability metrics, and end up writing a bad query that tries to divide a number by zero. It's become such a problem, that we think it would be best to make a specific type of error for division by zero.

Update the code so that the `divideError` type implements the `error` interface. Its `Error()` method should just return a string formatted in the following way:

```
can not divide DIVIDEND by zero
```

Where `DIVIDEND` is the actual dividend of the `divideError`. Use the `%v` [verb](https://pkg.go.dev/fmt#hdr-Printing) to format the dividend as a float.
