# Gator 🐊

Gator is a multi-user command-line RSS feed aggregator written in Go. It collects posts from RSS feeds, stores them in PostgreSQL, and lets you browse them right in your terminal.

## Features

- Register users and switch between them
- Add RSS feeds and follow or unfollow feeds other users have added
- Run a long-running aggregator that fetches feeds on an interval and saves new posts
- Browse the latest posts from the feeds you follow

## Requirements

You need these installed:

- [Go](https://go.dev/doc/install) 1.26 or later
- [PostgreSQL](https://www.postgresql.org/download/) 15 or later

## Installation

Install the `gator` CLI with `go install`:

```bash
go install github.com/vehspota/gator@latest
```

This puts a `gator` binary in `~/go/bin`. Make sure that folder is on your `PATH`.

## Setup

### 1. Create the database

Start PostgreSQL, then create a database called `gator`:

```bash
sudo -u postgres psql -c "CREATE DATABASE gator;"
```

### 2. Run the migrations

The schema is managed with [Goose](https://github.com/pressly/goose). Clone the repo and run the migrations:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
git clone https://github.com/vehspota/gator.git
cd gator/sql/schema
goose postgres "postgres://USER:PASSWORD@localhost:5432/gator" up
```

### 3. Create the config file

Create `~/.gatorconfig.json` in your home directory:

```json
{
  "db_url": "postgres://USER:PASSWORD@localhost:5432/gator?sslmode=disable"
}
```

Replace `USER` and `PASSWORD` with your PostgreSQL username and password. Gator fills in `current_user_name` for you.

## Usage

```bash
gator register alice
gator addfeed "Hacker News" "https://news.ycombinator.com/rss"
gator agg 1m
gator browse 5
```

### Commands

| Command | Description |
|---|---|
| `register <name>` | Create a new user and log in as them |
| `login <name>` | Switch to an existing user |
| `users` | List all users and show who is logged in |
| `addfeed <name> <url>` | Add a feed and follow it |
| `feeds` | List all feeds and who added them |
| `follow <url>` | Follow a feed another user added |
| `following` | List the feeds you follow |
| `unfollow <url>` | Stop following a feed |
| `agg <interval>` | Fetch feeds continuously, e.g. `30s`, `1m`, `1h` |
| `browse [limit]` | Show the newest posts from your feeds (default 2) |
| `reset` | Delete all users, feeds and posts. For development only! |

Run `agg` in one terminal and use the other commands in another.

## Built with

- Go, with `database/sql` and the `lib/pq` driver
- PostgreSQL
- [sqlc](https://sqlc.dev) for type-safe queries
- [Goose](https://github.com/pressly/goose) for migrations
