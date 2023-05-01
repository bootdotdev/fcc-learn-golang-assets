# Naming Generic Types

Let's look at this simple example again:

```go
func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
```

Remember, `T` is just a variable name, We could have named the type parameter *anything*. `T` happens to be a fairly common convention for a type variable, similar to how `i` is a convention for index variables in loops.

This is just as valid:

```go
func splitAnySlice[MyAnyType any](s []MyAnyType) ([]MyAnyType, []MyAnyType) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
```
