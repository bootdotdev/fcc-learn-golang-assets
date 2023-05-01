# Generics in Go

As we've mentioned, Go does *not* support classes. For a long time, that meant that Go code couldn't easily be reused in many circumstances. For example, imagine some code that splits a slice into 2 equal parts. The code that splits the slice doesn't really care about the *values* stored in the slice. Unfortunately in Go we would need to write it multiple times for each type, which is a very un-[DRY](https://blog.boot.dev/clean-code/dry-code/) thing to do.

```go
func splitIntSlice(s []int) ([]int, []int) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
```

```go
func splitStringSlice(s []string) ([]string, []string) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
```

In Go 1.20 however, support for [generics](https://blog.boot.dev/golang/how-to-use-golangs-generics/) was released, effectively solving this problem!

## Type Parameters

Put simply, generics allow us to use variables to refer to specific types. This is an amazing feature because it allows us to write abstract functions that drastically reduce code duplication.

```go
func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
```

In the example above, `T` is the name of the type parameter for the `splitAnySlice` function, and we've said that it must match the `any` constraint, which means it can be anything. This makes sense because the body of the function *doesn't care* about the types of things stored in the slice.

```go
firstInts, secondInts := splitAnySlice([]int{0, 1, 2, 3})
fmt.Println(firstInts, secondInts)
```

## Assignment

At Mailio we often store all the emails of a given email campaign in memory as a slice. We store payments for a single user in the same way.

Complete the `getLast()` function. It should be a generic function that returns the last element from a slice, no matter the types stored in the slice. If the slice is empty, it should return the zero value of the type.

## Tip: Zero value of a type

Creating a variable that's the zero value of a type is easy:

```go
var myZeroInt int
```

It's the same with generics, we just have a variable that represents the type:

```go
var myZero T
```
