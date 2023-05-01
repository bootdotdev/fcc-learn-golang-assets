# The Compilation Process

Computers need machine code, they don't understand English or even Go. We need to convert our high-level (Go) code into machine language, which is really just a set of instructions that some specific hardware can understand. In your case, your CPU.

The Go compiler's job is to take Go code and produce machine code. On Windows, that would be a `.exe` file. On Mac or Linux, it would be any executable file. The code you write in your browser here on Boot.dev is being compiled for you on Boot.dev's servers, then the machine code is executed in your browser as [Web Assembly](https://blog.boot.dev/golang/running-go-in-the-browser-with-web-assembly-wasm).

## A note on the structure of a Go program

We'll go over this all later in more detail, but to sate your curiosity for now, here are a few tidbits about the code.

1. `package main` lets the Go compiler know that we want this code to compile and run as a standalone program, as opposed to being a library that's imported by other programs.
2. `import fmt` imports the `fmt` (formatting) package. The formatting package exists in Go's standard library and let's us do things like print text to the console.
3. `func main()` defines the `main` function. `main` is the name of the function that acts as the entry point for a Go program.

## Assignment

To pass this exercise, fix the compiler error in the code.
