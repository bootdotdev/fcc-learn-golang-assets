# Passing Variables by Value

Variables in Go are passed by value (except for a few data types we haven't covered yet). "Pass by value" means that when a variable is passed into a function, that function receives a *copy* of the variable. The function is unable to mutate the caller's data.

```go
func main(){
    x := 5
    increment(x)

    fmt.Println(x)
    // still prints 5,
    // because the increment function
    // received a copy of x
}

func increment(x int){
    x++
}
```

## Assignment

It's critical in Textio that we keep track of how many SMS messages we send on behalf of our clients. Fix the bug to accurately track the number of SMS messages sent.

1. Alter the `incrementSends` function so that it returns the result after incrementing the `sendsSoFar`.
2. Alter `main()` to capture the return value from `incrementSends()` and overwrite the previous `sendsSoFar` value.
