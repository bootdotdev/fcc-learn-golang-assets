# Type Sizes

Ints, [uints](https://www.cs.utah.edu/~germain/PPS/Topics/unsigned_integer.html#:~:text=Unsigned%20Integers,negative%20(zero%20or%20positive).), [floats](https://techterms.com/definition/floatingpoint), and [complex](https://www.cloudhadoop.com/2018/12/golang-tutorials-complex-types-numbers.html#:~:text=Golang%20Complex%20Type%20Numbers,complex%20number%20is%2012.8i.) numbers all have type sizes.

```go
int  int8  int16  int32  int64 // whole numbers

uint uint8 uint16 uint32 uint64 uintptr // positive whole numbers

float32 float64 // decimal numbers

complex64 complex128 // imaginary numbers (rare)
```

The size (8, 16, 32, 64, 128, etc) indicates how many bits in memory will be used to store the variable. The default `int` and `uint` types are just aliases that refer to their respective 32 or 64 bit sizes depending on the environment of the user.

The standard sizes that should be used unless the developer has a specific need are:

* `int`
* `uint`
* `float64`
* `complex128`

Some types can be converted the following way:

```go
temperatureInt := 88
temperatureFloat := float64(temperatureInt)
```

Casting a float to an integer in this way [truncates](https://techterms.com/definition/truncate) the floating point portion.

## Assignment

Our Textio customers want to know how long they have had accounts with us.

Follow the instructions in the comment provided. You will create a new variable called `accountAgeInt` that will be the truncated, integer version of `accountAge`.
