# Range

Similar to slices and maps, channels can be ranged over.

```go
for item := range ch {
    // item is the next value received from the channel
}
```

This example will receive values over the channel (blocking at each iteration if nothing new is there) and will exit only when the channel is closed.

## Assignment

It's that time again, Mailio is hiring and we've been assigned to do the interview. For some reason, the [Fibonacci sequence](https://en.wikipedia.org/wiki/Fibonacci_number) is Mailio's interview problem of choice. We've been tasked with building a small toy program we can use in the interview.

Complete the `concurrrentFib` function. It should:

* Create a new channel of `int`s
* Call `fibonacci` in a goroutine, passing it the channel and the number of Fibonacci numbers to generate, `n`
* Use a `range` loop to read from the channel and print out the numbers one by one, each on a new line
