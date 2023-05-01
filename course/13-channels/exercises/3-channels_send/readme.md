# Channels

Empty structs are often used as `tokens` in Go programs. In this context, a token is a [unary](https://en.wikipedia.org/wiki/Unary_operation) value. In other words, we don't care *what* is passed through the channel. We care *when* and *if* it is passed.

We can block and wait until *something* is sent on a channel using the following syntax

```go
<-ch
```

This will block until it pops a single item off the channel, then continue, discarding the item.

## Assignment

Our Mailio server isn't able to boot up until it receives the signal that its databases are all online, and it learns about them being online by waiting for tokens (empty structs) on a channel.

Complete the `waitForDbs` function. It should block until it receives `numDBs` tokens on the `dbChan` channel. Each time it reads a token the `getDatabasesChannel` goroutine will print a message to the console for you.
