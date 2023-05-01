# Pointers

Pointers hold the memory address of a value.

The `*` syntax defines a pointer:

```go
var p *int
```

A pointer's zero value is `nil`

The & operator generates a pointer to its operand.

```go
myString := "hello"
myStringPtr = &myString
```

The * dereferences a pointer to gain access to the value

```go
fmt.Println(*myStringPtr) // read myString through the pointer
*myStringPtr = "world"    // set myString through the pointer 
```

Unlike C, Go has no [pointer arithmetic](https://www.tutorialspoint.com/cprogramming/c_pointer_arithmetic.htm)

## Just because you can doesn't mean you should

We're doing this exercise to understand that pointers **can** be used in this way. That said, pointers can be *very* dangerous. It's generally a better idea to have your functions accept non-pointers and return new values rather than mutating pointer inputs.

## Assignment

Complete the `removeProfanity` function.

It should use the [strings.ReplaceAll](https://pkg.go.dev/strings#ReplaceAll) function to replace all instances of the following words in the input `message` with asterisks.

* "dang" -> "****"
* "shoot" -> "*****"
* "heck" -> "****"

It should *mutate* the value in the pointer and return nothing.

Do *not* alter the function signature.
