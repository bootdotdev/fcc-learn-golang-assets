# Basic Types

Go's basic variable types are:

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

We talked about `string`s and `int`s previously, and those two types should be fairly self-explanatory.  A `bool` is a boolean variable, meaning it has a value of `true` or `false`.  The [floating point](https://en.wikipedia.org/wiki/Floating-point_arithmetic) types (`float32` and `float64`) are used for numbers that are not integers -- that is, they have digits to the right of the decimal place, such as `3.14159`.  The `float32` type uses 32 bits of precision, while the `float64` type uses 64 bits to be able to more precisely store more digits.  Don't worry too much about the intricacies of the other types for now.  We will cover some of them in more detail as the course progresses.

# Declaring a variable

Variables are declared using the `var` keyword. For example, to declare a variable called `number` of type `int`, you would write:

```go
var number int
```

To declare a variable called `pi` to be of type `float64` with a value of `3.14159`, you would write:

```go
var pi float64 = 3.14159
```

The value of an initialized variable with no assignment will be its [zero value](https://tour.golang.org/basics/12).

## Initialize some variables

Initialize the given variables to `int`, `float64`, `bool` and `string` respectively, with their zero values.
