# Feed Follows

Aside from just adding new feeds to the database, users can specify *which* feeds they want to follow. This will be important later when we want to show users a list of posts from the feeds they follow.

Add support for the following endpoints, and update the "create feed" endpoint as specified below.

## What is a "feed follow"?

A feed follow is just a link between a user and a feed. It's a [many-to-many](https://en.wikipedia.org/wiki/Many-to-many_(data_model)) relationship, so a user can follow many feeds, and a feed can be followed by many users.

Creating a feed follow indicates that a user is now following a feed. Deleting it is the same as "unfollowing" a feed.

It's important to understand that the `ID` of a feed follow is not the same as the `ID` of the feed itself. Each user/feed pair will have a unique feed follow id.

## Create a feed follow

Endpoint: `POST /v1/feed_follows`

*Requires authentication*

Example request body:

```json
{
  "feed_id": "4a82b372-b0e2-45e3-956a-b9b83358f86b"
}
```

Example response body:

```json
{
  "id": "c834c69e-ee26-4c63-a677-a977432f9cfa",
  "feed_id": "4a82b372-b0e2-45e3-956a-b9b83358f86b",
  "user_id": "0e4fecc6-1354-47b8-8336-2077b307b20e",
  "created_at": "2017-01-01T00:00:00Z",
  "updated_at": "2017-01-01T00:00:00Z"
}
```

## Delete a feed follow

Endpoint: `DELETE /v1/feed_follows/{feedFollowID}`

## Get all feed follows for a user

Endpoint: `GET /v1/feed_follows`

*Requires authentication*

Example response:

```json
[
  {
    "id": "c834c69e-ee26-4c63-a677-a977432f9cfa",
    "feed_id": "4a82b372-b0e2-45e3-956a-b9b83358f86b",
    "user_id": "0e4fecc6-1354-47b8-8336-2077b307b20e",
    "created_at": "2017-01-01T00:00:00Z",
    "updated_at": "2017-01-01T00:00:00Z"
  },
  {
    "id": "ad752167-f509-4ff3-8425-7781090b5c8f",
    "feed_id": "f71b842d-9fd1-4bc0-9913-dd96ba33bb15",
    "user_id": "0e4fecc6-1354-47b8-8336-2077b307b20e",
    "created_at": "2017-01-01T00:00:00Z",
    "updated_at": "2017-01-01T00:00:00Z"
  }
]
```

## Automatically create a feed follow when creating a feed

When a user creates a new feed, they should automatically be following that feed. They can of course choose to unfollow it later, but it should be there by default.

The response of this endpoint should now contain both entities:

```json
{
  "feed": { the feed object },
  "feed_follow": { the feed follow object }
}
```

## Test

As always, test all of your endpoints and make sure they work. Additionally, make sure that they return the proper error codes when they receive invalid inputs.
