# Constraints

Sometimes you need the logic in your generic function to know *something* about the types it operates on. The example we used in the first exercise didn't need to know *anything* about the types in the slice, so we used the built-in `any` constraint:

```go
func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
```

Constraints are just interfaces that allow us to write generics that only operate within the constraint of a given interface type. In the example above, the `any` constraint is the same as the empty interface because it means the type in question can be *anything*.

## Creating a custom constraint

Let's take a look at the example of a `concat` function. It takes a slice of values and concatenates the values into a string. This should work with *any type that can represent itself as a string*, even if it's not a string under the hood. For example, a `user` struct can have a `.String()` that returns a string with the user's name and age.

```go
type stringer interface {
    String() string
}

func concat[T stringer](vals []T) string {
    result := ""
    for _, val := range vals {
        // this is where the .String() method
        // is used. That's why we need a more specific
        // constraint instead of the any constraint
        result += val.String()
    }
    return result
}
```

## Assignment

We have different kinds of "line items" that we charge our customer's credit cards for. Line items can be things like "subscriptions" or "one-time payments" for email usage.

Complete the `chargeForLineItem` function. First, it should check if the user has a balance with enough funds to be able to pay for the cost of the `newItem`. If they don't then return an "insufficient funds" error.

If they *do* have enough funds:

* Add the line item to the user's history by appending the `newItem` to the slice of `oldItems`. This new slice is your first return value.
* Calculate the user's new balance by subtracting the cost of the new item from their balance. This is your second return value.
