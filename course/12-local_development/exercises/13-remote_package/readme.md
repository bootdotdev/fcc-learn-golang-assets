# Remote Packages

Let's learn how to use an open-source package that's available online.

## A note on how you should publish modules

Be aware that using the "replace" keyword like we did in the last assignment *isn't advised*, but can be useful to get up and running quickly. The *proper* way to create and depend on modules is to publish them to a remote repository. When you do that, the "replace keyword can be dropped from the `go.mod`:

### Bad

This works for local-only development

```go
module github.com/wagslane/hellogo

go 1.20

replace github.com/wagslane/mystrings v0.0.0 => ../mystrings

require (
	github.com/wagslane/mystrings v0.0.0
)
```

### Good

This works if we publish our modules to a remote location like Github as we should.

```go
module github.com/wagslane/hellogo

go 1.20

require (
	github.com/wagslane/mystrings v0.0.0
)
```

## Assignment

First, create a new directory in the same parent directory as `hellogo` and `mystrings` called `datetest`.

Create `main.go` in `datetest` and add the following code:

```go
package main

import (
	"fmt"
	"time"

	tinytime "github.com/wagslane/go-tinytime"
)

func main() {
	tt := tinytime.New(1585750374)

	tt = tt.Add(time.Hour * 48)
	fmt.Println(tt)
}
```

Initialize a module:

```bash
go mod init {REMOTE}/{USERNAME}/datetest
```

Download and install the remote go-tinydate package using `go get`:

```bash
go get github.com/wagslane/go-tinytime
```

Print the contents of your go.mod file to see the changes:

```bash
cat go.mod
```

Compile and run your program:

```bash
go build
./datetest
```
