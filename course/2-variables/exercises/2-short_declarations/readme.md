# Short Variable Declaration

Inside a function (even the main function), the `:=` short assignment statement can be used in place of a `var` declaration. The `:=` operator infers the type of the new variable based on the value.

```go
var empty string
```

Is the same as


```go
empty := ""
```

```go
numCars := 10 // inferred to be an integer

temperature := 0.0 // temperature is inferred to be a floating point value because it has a decimal point

var isFunny = true // isFunny is inferred to be a boolean
```

Outside of a function (in the [global/package scope](https://dave.cheney.net/2017/06/11/go-without-package-scoped-variables)), every statement begins with a keyword (`var`, `func`, and so on) and so the `:=` construct is not available.

## Assignment

A lot of our users send birthday messages using the Textio API.

Declare a variable named `congrats` with the value "happy birthday!" using a short variable declaration.
