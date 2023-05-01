# Go Run

Inside `hellogo`, create a new file called `main.go`.

Conventionally, the file in the `main` package that contains the `main()` function is called `main.go`.

Paste the following code into your file:

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

## Run the code

```bash
go run main.go
```

The `go run` command is used to quickly compile and run a Go package. The compiled binary is *not* saved in your working directory. Use `go build` instead to compile production executables.

I rarely use `go run` other than to quickly do some testing or debugging.

## Further reading

Execute `go help run` in your shell and read the instructions.
