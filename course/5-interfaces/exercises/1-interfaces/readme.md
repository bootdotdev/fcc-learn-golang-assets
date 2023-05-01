# Interfaces in Go

Interfaces are collections of method signatures. A type "implements" an interface if it has all of the methods of the given interface defined on it.

In the following example, a "shape" must be able to return its area and perimeter. Both `rect` and `circle` fulfill the interface.

```go
type shape interface {
  area() float64
  perimeter() float64
}

type rect struct {
    width, height float64
}
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perimeter() float64 {
    return 2*r.width + 2*r.height
}

type circle struct {
    radius float64
}
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}
```

When a type implements an interface, it can then be used as the interface type.

## Assignment

The `birthdayMessage` and `sendingReport` structs have already implemented the `getMessage` methods. The `getMessage` method simply returns a string, and any type that implements the method can be considered a `message`.

First, add the `getMessage()` method as a requirement on the method interface.

Second, complete the `sendMessage` function. It should print a message's `message`, which it obtains through the interface method. Notice that your code doesn't need to worry *at all* about whether a specific message is a `birthdayMessage` or a `sendingReport`!
