# Nested

Maps can contain maps, creating a nested structure. For example:

```go
map[string]map[string]int
map[rune]map[string]int
map[int]map[string]map[string]int
```

## Assignment

Because Textio is a glorified customer database, we have a lot of internal logic for sorting and dealing with customer names.

Complete the `getNameCounts` function. It takes a slice of strings (names) and returns a nested map where the first key is all the unique first characters of the names, the second key is all the names themselves, and the value is the count of each name.

For example:

```
billy
billy
bob
joe
```

Creates the following nested map:

```
b: {
    billy: 2,
    bob: 1
},
j: {
    joe: 1
}
```

Note that the test suite is *not* printing the map you're returning directly, but instead checking some specific keys.
