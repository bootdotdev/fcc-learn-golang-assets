# Maps

Maps are similar to JavaScript objects, Python dictionaries, and Ruby hashes. Maps are a data structure that provides key->value mapping.

The zero value of a map is `nil`.

We can create a map by using a literal or by using the `make()` function:

```go
ages := make(map[string]int)
ages["John"] = 37
ages["Mary"] = 24
ages["Mary"] = 21 // overwrites 24
```

```go
ages = map[string]int{
  "John": 37,
  "Mary": 21,
}
```

The `len()` function works on a map, it returns the total number of key/value pairs.

```go
ages = map[string]int{
  "John": 37,
  "Mary": 21,
}
fmt.Println(len(ages)) // 2
```

## Assignment

We can speed up our contact-info lookups by using a map! Looking up a value in a map by its key is much faster than searching through a slice.

Complete the `getUserMap` function. It takes a slice of names and a slice of phone numbers, and returns a map of `name` -> `user` structs and potentially an `error`. A `user` struct just contains a user's name and phone number.

If the length of `names` and `phoneNumbers` is not equal, return an error with the string "invalid sizes".

The first name in the `names` slice matches the first phone number, and so on.
