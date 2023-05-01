# Scraper

This is going to be a fairly large step. I recommend breaking it down into smaller pieces and functions, and testing each piece as you go.

Here are some different strategies I use depending on the situation:

* Write a unit test for a function that has simple inputs and outputs
* Edit `main.go` to call a function so I can quickly test it by running the whole program. Remove the call after testing and plug it into its proper place
* Put the code in a package, then write a separate `main` package (just a little `main()` script) that I can use to independently test the code in the package

*Commit your code each time you get a new piece working.*

## Add a `last_fetched_at` column to the `feeds` table

We need to keep track of when we last fetched the posts from a feed. This should be a nullable timestamp.

The `sql.NullTime` type is useful for nullable timestamps on the database side, but it's not great for marshaling into JSON. It results in a weird nested object. I'd recommend converting it to a `*time.Time` before returning it across the HTTP response.

I map all of my database structs to a different struct that has the intended JSON structure. This is a good way to keep your database and HTTP APIs separate.

For example: `func databaseFeedToFeed(feed database.Feed) Feed`

## Add `GetNextFeedsToFetch()` query to the database

It should return the next `n` feeds that need to be fetched, ordered by `last_fetched_at`, but with `NULL` values first. We obviously want to fetch the feeds that have never been fetched before or the ones that were fetched the longest time ago.

## Add a `MarkFeedFetched()` query to the database

It should update a feed and set its `last_fetched_at` to the current time. Don't forget to also update the `updated_at` field because we've updated the record.

## Write a function that can fetch data from a feed

This function should accept the URL of a live RSS feed, and return the parsed data in a Go struct.

You can test with these ones:

* `https://blog.boot.dev/index.xml`
* `https://wagslane.com/feed.xml`

And any other blogs you enjoy that have RSS feeds.

*Please be careful not to [DDOS](https://www.cloudflare.com/learning/ddos/what-is-a-ddos-attack/) any of the sites you're fetching from. Don't send too many requests!*

You can parse the returned XML with the [encoding/xml](https://pkg.go.dev/encoding/xml) package, it works *very* similarly to `encoding/json`. Define the structure of an RSS feed as a Go struct, then unmarshal the XML into that struct.

## Write a worker that fetches feeds continuously

This function should, on an interval (say every 60 seconds or so):

* Get the next `n` feeds to fetch from the database (you can configure `n`, I used `10`)
* Fetch and process all the feeds *at the same time* (you can use [sync.WaitGroup](https://pkg.go.dev/sync#WaitGroup) for this)

For now, "process" the feed by simply printing out the titles of each post

I recommend adding a lot of logging messages to this worker so that as it runs you can see what it's doing!

## Call your worker from `main.go`

Be sure to start the worker in its own goroutine, so that it runs in the background and processes feeds even as it simultaneously handles new HTTP requests.
