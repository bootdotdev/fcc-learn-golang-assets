# Errors Quiz

Go programs express errors with `error` values. Error-values are any type that implements the simple built-in [error interface](https://blog.golang.org/error-handling-and-go).

Keep in mind that the way Go handles errors is fairly unique. Most languages treat errors as something special and different. For example, Python raises exception types and JavaScript throws and catches errors. In Go, an `error` is just another value that we handle like any other value - however, we want! There aren't any special keywords for dealing with them.
