# Defer

The `defer` keyword is a fairly unique feature of Go. It allows a function to be executed automatically *just before* its enclosing function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

Deferred functions are typically used to close database connections, file handlers and the like.

For example:

```go
// CopyFile copies a file from srcName to dstName on the local filesystem.
func CopyFile(dstName, srcName string) (written int64, err error) {

  // Open the source file
  src, err := os.Open(srcName)
  if err != nil {
    return
  }
  // Close the source file when the CopyFile function returns
  defer src.Close()

  // Create the destination file
  dst, err := os.Create(dstName)
  if err != nil {
    return
  }
  // Close the destination file when the CopyFile function returns
  defer dst.Close()

  return io.Copy(dst, src)
}
```

In the above example, the `src.Close()` function is not called until after the `CopyFile` function was called but immediately before the `CopyFile` function returns.

Defer is a great way to **make sure** that something happens at the end of a function, even if there are multiple return statements.

## Assignment

There is a bug in the `logAndDelete` function, fix it!

This function should *always* delete the user from the user's map, which is a map that stores the user's name as keys. It also returns a `log` string that indicates to the caller some information about the user's deletion.

To avoid bugs like this in the future, instead of calling `delete` before each `return`, just `defer` the delete once at the beginning of the function.

