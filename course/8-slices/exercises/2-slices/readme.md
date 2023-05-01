# Slices in Go

*99 times out of 100* you will use a slice instead of an array when working with ordered lists.

Arrays are fixed in size. Once you make an array like `[10]int` you can't add an 11th element.

A slice is a *dynamically-sized*, *flexible* view of the elements of an array.

Slices **always** have an underlying array, though it isn't always specified explicitly. To explicitly create a slice on top of an array we can do:

```go
primes := [6]int{2, 3, 5, 7, 11, 13}
mySlice := primes[1:4]
// mySlice = {3, 5, 7}
```

The syntax is:

```
arrayname[lowIndex:highIndex]
arrayname[lowIndex:]
arrayname[:highIndex]
arrayname[:]
```

Where `lowIndex` is inclusive and `highIndex` is exclusive

Either `lowIndex` or `highIndex` or both can be omitted to use the entire array on that side.

## Assignment

Retries are a premium feature now! Textio's free users only get 1 retry message, while pro members get an unlimited amount.

Complete the `getMessageWithRetriesForPlan` function. It takes a `plan` variable as input, and you've been provided with constants for the plan types at the top of the program.

* If the plan is a pro plan, return all the strings from `getMessageWithRetries()`.
* If the plan is a free plan, return the first 2 strings from `getMessageWithRetries()`.
* If the plan isn't either of those, return an error that says `unsupported plan`.
