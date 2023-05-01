# Go is Strongly Typed

Go enforces strong and static typing, meaning variables can only have a single type. A `string` variable like "hello world" can not be changed to an `int`, such as the number `3`.

One of the biggest benefits of strong typing is that errors can be caught at "compile time". In other words, bugs are more easily caught ahead of time because they are detected when the code is compiled before it even runs.

Contrast this with most interpreted languages, where the variable types are dynamic. Dynamic typing can lead to subtle bugs that are hard to detect. With interpreted languages, the code *must* be run (sometimes in production if you are unlucky 😨) to catch syntax and type errors.

## Concatenating strings

Two strings can be [concatenated](https://en.wikipedia.org/wiki/Concatenation) with the `+` operator. Because Go is strongly typed, it won't allow you to concatenate a string variable with a numeric variable.

## Assignment

We'll be using simple [basic authentication](https://en.wikipedia.org/wiki/Basic_access_authentication) for the Textio API. This is how our users will communicate to us who they are and how many features they are paying for with each request to our API.

The code on the right has a type error. Change the type of `password` to a string (but use the same numeric value) so that it can be concatenated with the `username` variable.
