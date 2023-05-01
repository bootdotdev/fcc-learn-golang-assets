# Which Type Should I Use?

With so many types for what is essentially just a number, developers coming from languages that only have one kind of `Number` type (like JavaScript) may find the choices daunting.

## Prefer "default" types

A problem arises when we have a `uint16`, and the function we are trying to pass it into takes an `int`. We're forced to write code riddled with type casts like `int(myUint16)`.

This style of development can be slow and annoying to read. When Go developers stray from the “default” type for any given type family, the code can get messy quickly.

Unless you have a good reason to, stick to the following types:

* `bool`
* `string`
* `int`
* `uint32`
* `byte`
* `rune`
* `float64`
* `complex128`

## When should I use a more specific type?

When you're super concerned about performance and memory usage.

That’s about it. The only reason to deviate from the defaults is to squeeze out every last bit of performance when you are writing an application that is resource-constrained. (Or, in the special case of `uint64`, you need an absurd range of unsigned integers).

You can [read more on this subject here](https://blog.boot.dev/golang/default-native-types-golang/) if you'd like, but it's not required.
