# Go Install

## Build an executable

Ensure you are in your `hellogo` repo, then run:

```bash
go install
```

Navigate out of your project directory:

```bash
cd ../
```

Go has installed the `hellogo` program globally. Run it with:

```bash
hellogo
```

## Tip about "not found"

If you get an error regarding "hellogo not found" it means you probably don't have your Go environment setup properly. Specifically, `go install` is adding your binary to your `GOBIN` directory, but that may not be in your `PATH`.

You can read more about that here in the [go install docs](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies).
