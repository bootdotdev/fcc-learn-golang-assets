# Posts

## Add a `posts` table to the database

A post is a single entry from a feed. It should have:

* `id` - a unique identifier for the post
* `created_at` - the time the record was created
* `updated_at` - the time the record was last updated
* `title` - the title of the post
* `url` - the URL of the post *this should be unique*
* `description` - the description of the post
* `published_at` - the time the post was published
* `feed_id` - the ID of the feed that the post came from

Some of these fields can probably be null, others you might want to be more strict about - it's up to you.

## Add a "create post" SQL query to the database

This should insert a new post into the database.

## Add a "get posts by user" SQL query to the database

Order the results so that the most recent posts are first. Make the number of posts returned configurable.

## Update your scraper to save posts

Instead of just printing out the titles of the posts, save them to the database! If you encounter an error where the post with that URL already exists, just ignore it. That will happen a lot. If it's a different error, you should probably log it.

Make sure that you're parsing the "published at" time properly from the feeds. Sometimes they might be in a different format than you expect, so you might need to handle that.

## Add a "get posts by user" HTTP endpoint

Endpoint: `GET /v1/posts`

*This is an authenticated endpoint*

This endpoint should return a list of posts for the authenticated user. It should accept a `limit` query parameter that limits the number of posts returned. The default if the parameter is not provided can be whatever you think is reasonable.

## Start scraping some feeds!

Test your scraper to make sure it's working! Go find some of your favorite websites and add their RSS feeds to your database. Then start your scraper and watch it go to work.
