# Arrays in Go

Arrays are fixed-size groups of variables of the same type.

The type `[n]T` is an array of n values of type `T`

To declare an array of 10 integers:

```go
var myInts [10]int
```

or to declare an initialized literal:

```go
primes := [6]int{2, 3, 5, 7, 11, 13}
```

## Assignment

When a message isn't responded to, we allow our clients to have up to 2 additional messages that are sent as nudging reminders.

`getMessageWithRetries` returns an array of 3 strings where index `0` is the first message. If the first message isn't answered by the recipient, we send the second, if that one isn't answered then we send the third.

Update `getMessageWithRetries` to return the following 3 strings in an array.

* `click here to sign up`
* `pretty please click here`
* `we beg you to sign up`
