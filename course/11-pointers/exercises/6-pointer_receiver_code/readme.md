# Pointer Receiver Code

Methods with pointer receivers don't require that a pointer is used to call the method. The pointer will automatically be derived from the value.

```go
type circle struct {
	x int
	y int
    radius int
}

func (c *circle) grow(){
    c.radius *= 2
}

func main(){
    c := circle{
        x: 1,
        y: 2,
        radius: 4,
    }

    // notice c is not a pointer in the calling function
    // but the method still gains access to a pointer to c
    c.grow()
    fmt.Println(c.radius)
    // prints 8
}
```

## Assignment

Fix the bug in the code so that `setMessage` sets the `message` field of the given email structure, and the new value persists outside the scope of the `setMessage` method.
