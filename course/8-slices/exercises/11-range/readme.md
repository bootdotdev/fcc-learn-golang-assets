# Range

Go provides syntactic sugar to iterate easily over elements of a slice:

```go
for INDEX, ELEMENT := range SLICE {
}
```

For example:

```go
fruits := []string{"apple", "banana", "grape"}
for i, fruit := range fruits {
    fmt.Println(i, fruit)
}
// 0 apple
// 1 banana
// 2 grape
```

## Assignment

We need to be able to quickly detect bad words in the messages our system sends.

Complete the `indexOfFirstBadWord` function. If it finds any bad words in the message it should return the **index** of the first bad word in the `msg` slice. This will help us filter out naughty words from our messaging system. If no bad words are found, return `-1` instead.

Use the `range` keyword.
