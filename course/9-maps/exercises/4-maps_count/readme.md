# Count Instances

Remember that you can check if a key is already present in a map by using the second return value from the index operation.

```go
names := map[string]int{}

if _, ok := names["elon"]; !ok {
    // if the key doesn't exist yet,
    // initialize its value to 0
    names["elon"] = 0
}
```

## Assignment

We have a slice of user ids, and each instance of an id in the slice indicates that a message was sent to that user. We need to count up how many times each user's id appears in the slice to track how many messages they received.

Implement the `getCounts` function. It should return a map of `string -> int` so that each `int` is a count of how many times each string was found in the slice.
