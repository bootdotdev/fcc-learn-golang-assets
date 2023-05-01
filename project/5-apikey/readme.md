# API Key

## 1. Add an "api key" column to the users table

Use a new migration file in the `sql/schema` directory to add a new column to the `users` table. I named my file `002_users_apikey.sql`.

The "up" migration adds the column, and the "down" migration removes it.

Use a `VARCHAR(64)` that must be unique and not null. Using a string of a specific length does two things:

1. It ensures we don't accidentally store a key that's too long (type safety)
2. It's more performant than using a variable length `TEXT` column

Because we're enforcing the `NOT NULL` constraint, and we already have some users in the database, we need to set a default value for the column. A blank default would be a bit silly: that's no better than null! Instead, we'll generate valid API keys (256-bit hex values) using SQL. Here's the function I used:

```sql
encode(sha256(random()::text::bytea), 'hex')
```

When you're done, use `goose postgres CONN up` to perform the migration.

## 2. Create an API key for new users

Update your "create user" SQL query to use the same SQL function to generate API keys for new users.

## 3. Add a new SQL query to get a user by their API key

This query can live in the same file as the "create user" query, or you can make a new one - it's up to you.

## 4. Generate new Go code

Run `sqlc generate` to generate new Go code for your queries.

## 5. New endpoint: return the current user

Add a new endpoint that allows users to get their own user information. You'll need to parse the header and use your new query to get the user data.

Endpoint: `GET /v1/users`

Request headers: `Authorization: ApiKey <key>`

Example response body:

```json
{
    "id": "3f8805e3-634c-49dd-a347-ab36479f3f83",
    "created_at": "2021-09-01T00:00:00Z",
    "updated_at": "2021-09-01T00:00:00Z",
    "name": "Lane",
    "api_key": "cca9688383ceaa25bd605575ac9700da94422aa397ef87e765c8df4438bc9942"
}
```

*Test your endpoints with an HTTP client before moving on!*

Don't forget that each time you update your queries or schema you'll need to regenerate your Go code with `sqlc generate`. If you update the schema you'll also need to migrate your database up (and maybe down).
