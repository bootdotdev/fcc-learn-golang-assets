# Interfaces Quiz

Remember, interfaces are collections of method signatures. A type "implements" an interface if it has all of the methods of the given interface defined on it.

```go
type shape interface {
  area() float64
}
```

If a type in your code implements an `area` method, with the same signature (e.g. accepts nothing and returns a `float64`), then that object is said to *implement* the `shape` interface.

```go
type circle struct{
	radius int
}

func (c *circle) area() float64 {
  return 3.14 * c.radius * c.radius
}
```

This is *different from most other languages*, where you have to *explicitly* assign an interface type to an object, like with Java:

```java
class Circle implements Shape
```
