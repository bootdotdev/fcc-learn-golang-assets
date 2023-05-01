# Structs in Go

We use structs in Go to represent structured data. It's often convenient to group different types of variables together. For example, if we want to represent a car we could do the following:

```go
type car struct {
  Make string
  Model string
  Height int
  Width int
}
```

This creates a new struct type called `car`. All cars have a `Make`, `Model`, `Height` and `Width`.

In Go, you will often use a struct to represent information that you would have used a dictionary for in Python, or an object literal for in JavaScript.

## Assignment

Complete the `messageToSend` struct definition. It needs two fields:

* `phoneNumber` - an integer
* `message` - a string.
