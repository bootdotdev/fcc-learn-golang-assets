# Tricky Slices

The `append()` function changes the underlying array of its parameter AND returns a new slice. This means that using `append()` on anything other than itself is usually a BAD idea.

```go
// dont do this!
someSlice = append(otherSlice, element)
```

Take a look at these head-scratchers:

## Example 1: Works as expected

```go
a := make([]int, 3)
fmt.Println("len of a:", len(a))
// len of a: 3 
fmt.Println("cap of a:", cap(a))
// cap of a: 3
fmt.Println("appending 4 to b from a")
// appending 4 to b from a
b := append(a, 4)
fmt.Println("b:", b)
// b: [0 0 0 4]
fmt.Println("addr of b:", &b[0])
// addr of b: 0x44a0c0
fmt.Println("appending 5 to c from a")
// appending 5 to c from a
c := append(a, 5)
fmt.Println("addr of c:", &c[0])
// addr of c: 0x44a180
fmt.Println("a:", a)
// a: [0 0 0]
fmt.Println("b:", b)
// b: [0 0 0 4]
fmt.Println("c:", c)
// c: [0 0 0 5]
```

With slices `a`, `b`, and `c`, `4` and `5` seem to be appended as we would expect. We can even check the memory addresses and confirm that `b` and `c` point to different underlying arrays.

## Example 2: Something fishy

```go
i := make([]int, 3, 8)
fmt.Println("len of i:", len(i))
// len of i: 3
fmt.Println("cap of i:", cap(i))
// cap of i: 8
fmt.Println("appending 4 to j from i")
// appending 4 to j from i
j := append(i, 4)
fmt.Println("j:", j)
// j: [0 0 0 4]
fmt.Println("addr of j:", &j[0])
// addr of j: 0x454000
fmt.Println("appending 5 to g from i")
// appending 5 to g from i
g := append(i, 5)
fmt.Println("addr of g:", &g[0])
// addr of g: 0x454000
fmt.Println("i:", i)
// i: [0 0 0]
fmt.Println("j:", j)
// j: [0 0 0 5]
fmt.Println("g:", g)
// g: [0 0 0 5]
```

In this example however, when `5` is appended to `g` it overwrites `j`'s fourth index because `j` and `g` point to the *same underlying array*. The `append()` function only creates a new array when there isn't any capacity left. We created `i` with a length of 3 and a capactiy of 8, which means we can append `5` items before a new array is automatically allocated.

Again, to avoid bugs like this, you should always use the `append` function on the same slice the result is assigned to:

```go
mySlice := []int{1, 2, 3}
mySlice = append(mySlice, 4)
```
