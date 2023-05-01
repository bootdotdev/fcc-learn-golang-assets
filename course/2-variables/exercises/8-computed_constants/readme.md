# Computed Constants

Constants must be known at compile time. More often than not they will be declared with a static value:

```go
const myInt = 15
```

However, constants *can be computed* so long as the computation can happen at *compile time*.

For example, this is valid:

```go
const firstName = "Lane"
const lastName = "Wagner"
const fullName = firstName + " " + lastName
```

That said, you *cannot* declare a constant that can only be computed at run-time.

## Assignment

Keeping track of time in a message-sending application is *critical*. Imagine getting an appointment reminder an hour **after** your doctor's visit.

Complete the code using a computed constant to print the number of seconds in an hour.
