# There is no while loop in Go

Most programming languages have a concept of a `while` loop. Because Go allows for the omission of sections of a `for` loop, a `while` loop is just a `for` loop that only has a CONDITION.

```go
for CONDITION {
  // do some stuff while CONDITION is true
}
```

For example:

```go
plantHeight := 1
for plantHeight < 5 {
  fmt.Println("still growing! current height:", plantHeight)
  plantHeight++
}
fmt.Println("plant has grown to ", plantHeight, "inches")
```

Which prints:

```
still growing! current height: 1
still growing! current height: 2
still growing! current height: 3
still growing! current height: 4
plant has grown to 5 inches
```

## Assignment

We have an interesting new cost structure from our SMS vendor. They charge exponentially more money for each consecutive text we send! Let's write a function that can calculate how many messages we can send in a given batch given a `costMultiplier` and a `maxCostInPennies`.

In a nutshell, the first message costs a penny, and each message after that costs the same as the previous message multiplied by the `costMultiplier`. That gets expensive!

There is an infinite loop in the code! Let's add a condition to fix the bug. The loop should exit **before** incrementing `maxMessagesToSend` if the cost of the next message would go over the max cost.
