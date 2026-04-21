# bloggator

A CLI RSS feed aggregator written in Go. Maybe for blogs and gator as in "aggregator", hence bloggator.

## Prerequisites

- Go installed
- PostgreSQL installed

## Installation

```bash
go install github.com/Puhkusarvikuono/bloggator@latest
```

## Config

Config keeps track of logged in user and connection credentials for PostgreSQL database.
You will need a config file ~/.gatorconfig.json with database connection string. 
Replace example url with a real Postgres connection string:

```json
{
    "db_url": "protocol://username:password@host:port/database",
    "current_user_name": ""
}
```

## Using bloggator

Bloggator is a multi user CLI application with several commands.

All commands are run as:

```bash
bloggator <command> [args]
```

### User management

| Command | Arguments | Description |
|---|---|---|
| `register` | `<username>` | Creates a new user and logs them in. |
| `login` | `<username>` | Logs in an existing user. |
| `users` | — | Lists all registered users and marks the current one. |
| `reset` | — | Wipes all users and data from the database. Use with caution. |

### Feeds

| Command | Arguments | Description |
|---|---|---|
| `addfeed`* | `<name> <url>` | Adds a new RSS feed and auto-follows it. |
| `feeds` | — | Lists every feed in the database along with the user who added it. |
| `agg` | `<time_between_reqs>` | Starts the long-running aggregator that fetches feeds on an interval (e.g. `1m`, `30s`). |

### Following & browsing

| Command | Arguments | Description |
|---|---|---|
| `follow`* | `<url>` | Follows an existing feed by URL. |
| `following`* | — | Lists the feeds the current user follows. |
| `unfollow`* | `<url>` | Unfollows a feed by URL. |
| `browse`* | `[limit]` | Shows the most recent posts from feeds you follow (default limit: 2). |

\* Requires being logged in.

### Examples

```bash
bloggator register lane
bloggator addfeed "Boot.dev Blog" https://blog.boot.dev/index.xml
bloggator agg 1m
bloggator browse 5
```


