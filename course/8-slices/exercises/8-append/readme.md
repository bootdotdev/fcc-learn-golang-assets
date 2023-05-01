# Append

The built-in append function is used to dynamically add elements to a slice:

```go
func append(slice []Type, elems ...Type) []Type
```

If the underlying array is not large enough, `append()` will create a new underlying array and point the slice to it.

Notice that `append()` is variadic, the following are all valid:

```go
slice = append(slice, oneThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice...)
```

## Assignment

We've been asked to "bucket" costs for an entire month into the cost occurring on each day of the month.

Complete the `getCostsByDay` function. It should return a slice of `float64`s, where each element is the total cost for that day. The length of the slice should be equal to the number of days represented in the `costs` slice, including any days that have no costs, up to the last day represented in the slice.

Days are numbered and start at `0`.

### Example

Input in `day, cost` format:

```go
[]cost{
    {0, 4.0},
    {1, 2.1},
    {1, 3.1},
    {5, 2.5},
}
```

I would expect this result:

```go
[]float64{
    4.0,
    5.2,
    0.0,
    0.0,
    0.0,
    2.5,
}
```
