# Select

Sometimes we have a single goroutine listening to multiple channels and want to process data in the order it comes through each channel.

A `select` statement is used to listen to multiple channels at the same time. It is similar to a `switch` statement but for channels.

```go
select {
  case i, ok := <- chInts:
    fmt.Println(i)
  case s, ok := <- chStrings:
    fmt.Println(s)
}
```

The first channel with a value ready to be received will fire and its body will execute. If multiple channels are ready at the same time one is chosen randomly. The `ok` variable in the example above refers to whether or not the channel has been closed by the sender yet.

## Assignment

Complete the `logMessages` function.

Use an infinite `for` loop and a select statement to log the emails and sms messages as they come in order across the two channels. Add a condition to `return` from the function when *one* of the two channels closes, whichever is first.

Use the `logSms` and `logEmail` functions to log the messages.
