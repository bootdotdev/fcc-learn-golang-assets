# Boilerplate

Before we get to the app-specific stuff, let's scaffold a simple CRUD server, hopefully, you're already familiar with how to do this from the "Learn Web Servers" course! That said, I'll provide a quick refresher.

*It might be a good idea to use your "Learn Web Servers" code as a reference while building this project!*

## 1. Create a new project

You should know how to do this by now! My process is:

* Create a repo on GitHub or GitLab (initialized with a README)
* Clone it onto your machine
* Create a new Go module with `go mod init`.
* Create a `main.go` file in the root of your project, and add a `func main()` to it.

## 2. Install packages

Install the following packages using `go get`:

* [chi](https://github.com/go-chi/chi)
* [cors]https://github.com/go-chi/cors
* [godotenv](github.com/joho/godotenv)

## 3. Env

Create a gitignore'd `.env` file in the root of your project and add the following:

```bash
PORT="8080"
```

The `.env` file is a convenient way to store environment (configuration) variables.

* Use [godotenv.Load()](https://pkg.go.dev/github.com/joho/godotenv#Load) to load the variables from the file into your environment at the top of `main()`.
* Use [os.Getenv()](https://pkg.go.dev/os#Getenv) to get the value of `PORT`.

## 4. Create a router and server

1. Create a [chi.NewRouter](https://pkg.go.dev/github.com/go-chi/chi#NewRouter)
2. Use [router.Use](https://pkg.go.dev/github.com/go-chi/chi#Router.Use) to add the built-in [cors.Handler](https://pkg.go.dev/github.com/go-chi/cors#Handler) middleware.
3. Create sub-router for the `/v1` namespace and mount it to the main router.
4. Create a new [http.Server](https://pkg.go.dev/net/http#Server) and add the port and the main router to it.
5. Start the server

## 5. Create some JSON helper functions

Create two functions:

* `respondWithJSON(w http.ResponseWriter, status int, payload interface{})`
* `respondWithError(w http.ResponseWriter, code int, msg string)` (which calls `respondWithJSON` with error-specific values)

You used these in the "Learn Web Servers" course, so you should be able to figure out how to implement them again. They're simply helper functions that write an HTTP response with:

* A status code
* An `application/json` content type
* A JSON body

## 6. Add a readiness handler

Add a handler for `GET /v1/readiness` requests. It should return a 200 status code and a JSON body:

```json
{
  "status": "ok"
}
```

*The purpose of this endpoint is for you to test your `respondWithJSON` function.*

## 7. Add an error handler

Add a handler for `GET /v1/err` requests. It should return a 500 status code and a JSON body:

```json
{
  "error": "Internal Server Error"
}
```

*The purpose of this endpoint is for you to test your `respondWithError` function.*

## 8. Run and test your server

```bash
go build -o out && ./out
```

Once it's running, use an HTTP client to test your endpoints.
