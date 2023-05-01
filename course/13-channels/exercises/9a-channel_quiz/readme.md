# Channels Review

Here are a few extra things you should understand about channels from [Dave Cheney's awesome article](https://dave.cheney.net/2014/03/19/channel-axioms).

## A send to a nil channel blocks forever

```go
var c chan string // c is nil
c <- "let's get started" // blocks
```

## A receive from a nil channel blocks forever

```go
var c chan string // c is nil
fmt.Println(<-c) // blocks
```

## A send to a closed channel panics

```go
var c = make(chan int, 100)
close(c)
c <- 1 // panic: send on closed channel
```

## A receive from a closed channel returns the zero value immediately

```go
var c = make(chan int, 100)
close(c)
fmt.Println(<-c) // 0
```
