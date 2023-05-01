# Create a Feed

An RSS feed is just a URL that points to some XML. Users will be able to add feeds to our database so that our server (in a future step) can go download all of the posts in the feed (like blog posts or podcast episodes).

## 1. Create a feeds table

Like any table in our DB, we'll need the standard `id`, `created_at`, and `updated_at` fields. We'll also need a few more:

* `name`: The name of the feed (like "The Changelog, or "The Boot.dev Blog")
* `url`: The URL of the feed
* `user_id`: The ID of the user who added this feed

I'd recommend making the `url` field unique so that in the future we aren't downloading duplicate posts. I'd also recommend using [ON DELETE CASCADE](https://stackoverflow.com/a/14141354) on the `user_id` foreign key so that if a user is deleted, all of their feeds are automatically deleted as well.

Write the appropriate migrations and run them.

## 2. Add a query to create a feed

Add a new query to create a feed, then use `sqlc generate` to generate the Go code.

## 3. Create some authentication middleware

Most of the endpoints going forward will require a user to be logged in. Let's DRY up our code by creating some middleware that will check for a valid API key.

Now, I'm not a fan of how the Chi router handles stateful middleware using [context](https://pkg.go.dev/context) (middleware that passes data down to the next handler). I prefer to create custom handlers that accept extra values. You can add middleware however you like, but here are some examples from my code.

### A custom type for handlers that require authentication

```go
type authedHandler func(http.ResponseWriter, *http.Request, database.User)
```

### Middleware that authenticates a request, gets the user and calls the next authed handler

```go
func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
    ///
}
```

### Using the middleware

```go
v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerUsersGet))
```

## 4. Create a handler to create a feed

Create a handler that creates a feed. This handler *and* the "get user" handler should use the authentication middleware.

Endpoint: `POST /v1/feeds`

Example request body:

```json
{
  "name": "The Boot.dev Blog",
  "url": "https://blog.boot.dev/index.xml"
}
```

Example response body:

```json
{
  "id": "4a82b372-b0e2-45e3-956a-b9b83358f86b",
  "created_at": "2021-05-01T00:00:00Z",
  "updated_at": "2021-05-01T00:00:00Z",
  "name": "The Boot.dev Blog",
  "url": "https://blog.boot.dev/index.xml",
  "user_id": "d6962597-f316-4306-a929-fe8c8651671e"
}
```

## 5. Test

Test your handler using an HTTP client, then use your database client to make sure the data was saved correctly.
