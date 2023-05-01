# Multiple Interfaces

A type can implement any number of interfaces in Go. For example, the empty interface, `interface{}`, is *always* implemented by every type because it has no requirements.

## Assignment

Add the required methods so that the `email` type implements both the `expense` and `printer` interfaces.

### cost()

If the email is *not* "subscribed", then the cost is `0.05` for each character in the body. If it *is*, then the cost is `0.01` per character.

### print()

The `print` method should print to standard out the email's body text.
