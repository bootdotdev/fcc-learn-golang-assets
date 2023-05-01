# Create Users

In this step, we'll be adding an endpoint to create new users on the server. We'll be using a couple of tools to help us out:

* [database/sql](https://pkg.go.dev/database/sql): This is part of Go's standard library. It provides a way to connect to a SQL database, execute queries, and scan the results into Go types.
* [sqlc](https://sqlc.dev/): SQLC is an *amazing* Go program that generates Go code from SQL queries. It's not exactly an [ORM](https://www.freecodecamp.org/news/what-is-an-orm-the-meaning-of-object-relational-mapping-database-tools/), but rather a tool that makes working with raw SQL almost as easy as using an ORM.
* [Goose](https://github.com/pressly/goose): Goose is a database migration tool written in Go. It runs migrations from the same SQL files that SQLC uses, making the pair of tools a perfect fit.

## 1. Install SQLC

SQLC is just a command line tool, it's not a package that we need to import. I recommend [installing](https://docs.sqlc.dev/en/latest/overview/install.html) it using `go install`:

```bash
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

Then run `sqlc version` to make sure it's installed correctly.

## 2. Install Goose

Like SQLC, Goose is just a command line tool. I also recommend [installing](https://github.com/pressly/goose#install) it using `go install`:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Run `goose -version` to make sure it's installed correctly.

## 3. Create the `users` migration

I recommend creating an `sql` directory in the root of your project, and in there creating a `schema` directory.

A "migration" is a SQL file that describes a change to your database schema. For now, we need our first migration to create a `users` table. The simplest format for these files is:

```
number_name.sql
```

For example, I created a file in `sql/schema` called `001_users.sql` with the following contents:

```sql
-- +goose Up
CREATE TABLE ...

-- +goose Down
DROP TABLE users;
```

Write out the `CREATE TABLE` statement in full, I left it blank for you to fill in. A `user` should have 4 fields:

* id: a `UUID` that will serve as the primary key
* created_at: a `TIMESTAMP` that can not be null
* updated_at: a `TIMESTAMP` that can not be null
* name: a string that can not be null

The `-- +goose Up` and `-- +goose Down` comments are required. They tell Goose how to run the migration. An "up" migration moves your database from its old state to a new state. A "down" migration moves your database from its new state back to its old state.

By running all of the "up" migrations on a blank database, you should end up with a database in a ready-to-use state. "Down" migrations are only used when you need to roll back a migration, or if you need to reset a local testing database to a known state.

## 4. Run the migration

`cd` into the `sql/schema` directory and run:

```bash
goose postgres CONN up
```

Where `CONN` is the connection string for your database. Here is mine:

```
postgres://wagslane:@localhost:5432/blogator
```

The format is:

```
protocol://username:password@host:port/database
```

Run your migration! Make sure it works by using PGAdmin to find your newly created `users` table.

## 5. Save your connection string as an environment variable

Add your connection string to your `.env` file. When using it with `goose`, you'll use it in the format we just used. However, here in the `.env` file it needs an additional query string:

```
protocol://username:password@host:port/database?sslmode=disable
```

Your application code needs to know to not try to use SSL locally.

## 6. Configure [SQLC](https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html)

You'll always run the `sqlc` command from the root of your project. Create a file called `sqlc.yaml` in the root of your project. Here is mine:

```yaml
version: "2"
sql:
  - schema: "sql/schema"
    queries: "sql/queries"
    engine: "postgresql"
    gen:
      go:
        out: "internal/database"
```

We're telling SQLC to look in the `sql/schema` directory for our schema structure (which is the same set of files that Goose uses, but sqlc automatically ignores "down" migrations), and in the `sql/queries` directory for queries. We're also telling it to generate Go code in the `internal/database` directory.

## 7. Write a query to create a user

Inside the `sql/queries` directory, create a file called `users.sql`. Here is mine:

```sql
-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;
```

`$1`, `$2`, `$3`, and `$4` are parameters that we'll be able to pass into the query in our Go code. The `:one` at the end of the query name tells SQLC that we expect to get back a single row (the created user).

Keep the [SQLC docs](https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html) handy, you'll probably need to refer to them again later.

## 8. Generate the Go code

Run `sqlc generate` from the root of your project. It should create a new package of go code in `internal/database`.

## 9. Open a connection to the database, and store it in a config struct

If you recall from the web servers project, it's common to use a "config" struct to store shared data that HTTP handlers need access to. We'll do the same thing here. Mine looks like this:

```go
type apiConfig struct {
	DB *database.Queries
}
```

At the top of `main()` load in your database URL from your `.env` file, and then [.Open()](https://pkg.go.dev/database/sql#Open) a connection to your database:

```go
db, err := sql.Open("postgres", dbURL)
```

Use your generated `database` package to create a new `*database.Queries`, and store it in your config struct:

```go
dbQueries := database.New(db)
```

## 10. Create an HTTP handler to create a user

Endpoint: `POST /v1/users`

Example body:

```json
{
  "name": "Lane"
}
```

Example response:

```json
{
    "id": "3f8805e3-634c-49dd-a347-ab36479f3f83",
    "created_at": "2021-09-01T00:00:00Z",
    "updated_at": "2021-09-01T00:00:00Z",
    "name": "Lane"
}
```

Use Google's [UUID](https://pkg.go.dev/github.com/google/uuid) package to generate a new [UUID](https://blog.boot.dev/clean-code/what-are-uuids-and-should-you-use-them/) for the user's ID. Both `created_at` and `updated_at` should be set to the current time. If we ever need to update a user, we'll update the `updated_at` field.

I'm a fan of a convention where *every table* in my database has:

* An `id` field that is a UUID (if you're curious why, [read this](https://blog.boot.dev/clean-code/what-are-uuids-and-should-you-use-them/))
* A `created_at` field that indicates when the row was created
* An `updated_at` field that indicates when the row was last updated

## 11. Test your handler with an HTTP client!

C'mon, you know what to do.
