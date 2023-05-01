# Omitting conditions from a for loop in Go

Loops in Go can omit sections of a for loop. For example, the `CONDITION` (middle part) can be omitted which causes the loop to run forever.

```go
for INITIAL; ; AFTER {
  // do something forever
}
```

## Assignment

Complete the `maxMessages` function. Given a cost threshold, it should calculate the maximum number of messages that can be sent.

Each message costs `1.0`, plus an additional fee. The fee structure is:

* 1st message: `1.0 + 0.00`
* 2nd message: `1.0 + 0.01`
* 3rd message: `1.0 + 0.02`
* 4th message: `1.0 + 0.03`

## Browser freeze

If you lock up your browser by creating an infinite loop that isn't breaking, just click the `cancel` button.
