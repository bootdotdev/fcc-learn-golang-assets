# Nested structs in Go

Structs can be nested to represent more complex entities:

```go
type car struct {
  Make string
  Model string
  Height int
  Width int
  FrontWheel Wheel
  BackWheel Wheel
}

type Wheel struct {
  Radius int
  Material string
}
```

The fields of a struct can be accessed using the dot `.` operator.

```go
myCar := car{}
myCar.FrontWheel.Radius = 5
```

## Assignment

Textio has a bug, we've been sending texts with information missing! Before we send text messages in Textio, we should check to make sure the required fields have non-zero values.

Notice that the `user` struct is a nested struct within the `messageToSend` struct. Both `sender` and `recipient` are `user` struct types.

Complete the `canSendMessage` function. It should return `true` only if the `sender` and `recipient` fields each contain a `name` and a `number`. If any of the default zero values are present, return `false` instead.
