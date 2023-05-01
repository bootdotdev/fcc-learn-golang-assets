# Name Your Interface Arguments

Consider the following interface:

```go
type Copier interface {
  Copy(string, string) int
}
```

Based on the code alone, can you deduce what *kinds* of strings you should pass into the `Copy` function?

We know the function signature expects 2 string types, but what are they? Filenames? URLs? Raw string data? For that matter, what the heck is that `int` that's being returned?

Let's add some named arguments and return data to make it more clear.

```go
type Copier interface {
  Copy(sourceFile string, destinationFile string) (bytesCopied int)
}
```

Much better. We can see what the expectations are now. The first argument is the `sourceFile`, the second argument is the `destinationFile`, and `bytesCopied`, an integer, is returned.
