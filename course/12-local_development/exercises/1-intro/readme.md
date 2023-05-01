# Packages

Every Go program is made up of packages.

You have probably noticed the `package main` at the top of all the programs you have been writing.

A package named "main" has an entrypoint at the `main()` function. A `main` package is compiled into an executable program.

A package by any other name is a "library package". Libraries have no entry point. Libraries simply export functionality that can be used by other packages. For example:

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

This program is an executable. It is a "main" package and *imports* from the `fmt` and `math/rand` library packages.

## Assignment

Fix the bug in the code.
