# Type Inference

To declare a variable without specifying an explicit type (either by using the `:=` syntax or `var = expression` syntax), the variable's type is *inferred* from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type:

```go
var i int
j := i // j is also an int
```

However, when the right hand side is a literal value (an untyped numeric constant like `42` or `3.14`), the new variable will be an `int`, `float64`, or `complex128` depending on its precision:

```go
i := 42           // int
f := 3.14         // float64
g := 0.867 + 0.5i // complex128
```

## Assignment

Our current price to send a text message is 2 cents. However, it's likely that in the future the price will be a fraction of a penny, so we should use a `float64` to store this value.

Edit the `penniesPerText` declaration so that it's inferred by the compiler to be a `float64`.
