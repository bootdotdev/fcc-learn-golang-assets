# The Errors Package

The Go standard library provides an "errors" package that makes it easy to deal with errors.

Read the godoc for the [errors.New()](https://pkg.go.dev/errors#New) function, but here's a simple example:

```go
var err error = errors.New("something went wrong")
```

## Assignment

Twilio's software architects may have overcomplicated the requirements from the last coding assignment... oops. All we needed was a new generic error message that returns the string `no dividing by 0` when a user attempts to get us to perform the taboo.

Complete the `divide` function. Use the `errors.New()` function to return an error when `y == 0` that reads "no dividing by 0".
