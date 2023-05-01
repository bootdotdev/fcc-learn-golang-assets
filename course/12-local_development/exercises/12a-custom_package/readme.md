# Custom Package Continued

Let's use our new `mystrings` package in `hellogo`

Modify hellogo's `main.go` file:

```go
package main

import (
	"fmt"

	"{REMOTE}/{USERNAME}/mystrings"
)

func main() {
	fmt.Println(mystrings.Reverse("hello world"))
}
```

Don't forget to replace {REMOTE} and {USERNAME} with the values you used before. Then edit hellogo's `go.mod` file to contain the following:

```go
module example.com/username/hellogo

go 1.20

replace example.com/username/mystrings v0.0.0 => ../mystrings

require (
	example.com/username/mystrings v0.0.0
)
```

Now build and run the new program:

```bash
go build
./hellogo
```
