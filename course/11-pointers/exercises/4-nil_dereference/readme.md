# Nil Pointers

Pointers can be very dangerous.

If a pointer points to nothing (the zero value of the pointer type) then dereferencing it will cause a runtime error (a [panic](https://gobyexample.com/panic)) that crashes the program. Generally speaking, whenever you're dealing with pointers you should check if it's `nil` before trying to dereference it.

## Assignment

Let's make our profanity checker *safe*. Update the `removeProfanity` function. If `message` is `nil`, `return` early to avoid a [panic](https://gobyexample.com/panic). After all, there are no bad words to remove.
