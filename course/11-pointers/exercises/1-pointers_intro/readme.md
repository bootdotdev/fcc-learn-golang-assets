# Introduction to Pointers

As we have learned, a variable is a named location in memory that stores a value. We can manipulate the value of a variable by assigning a new value to it or by performing operations on it. When we assign a value to a variable, we are storing that value in a specific location in memory.

```go
x := 42
// "x" is the name of a location in memory. That location is storing the integer value of 42
```

## A pointer is a variable

A pointer is a variable that stores the *memory address* of another variable. This means that a pointer "points to" the *location* of where the data is stored *NOT* the actual data itself.

The `*` syntax defines a pointer:

```go
var p *int
```

The `&` operator generates a pointer to its operand.

```go
myString := "hello"
myStringPtr = &myString
```

## Why are pointers useful?

Pointers allow us to manipulate data in memory directly, without making copies or duplicating data. This can make programs more efficient and allow us to do things that would be difficult or impossible without them.

## Assignment

Fix the bug in the `sendMessage` function. It's *supposed* to print a nicely formatted message to the console containing an SMS's recipient and message body. However, it's not working as expected. Run the code and see what happens, then fix the bug.
