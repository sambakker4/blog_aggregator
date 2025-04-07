# Blog Aggregator Gator

In order for this CLI application to work, you'll need to first install Postgres and Go

## Install Postgres
### With macOS
`brew install postgresql@15`

### With Linux / WSL Debian
`sudo apt update
sudo apt install postgresql postgresql-contrib`

### Run to ensure installed
`psql --version`

### (Linux only) Update postgres password
`sudo passwd postgres`

## Start the Postgres server
### With macOS
`brew services start postgresql@15`

### With Linux
`sudo service postgresql start`

## Install Go
### Linux
`sudo apt install golang-go`
### With macOS
`brew install go`
### Insure installed
`go version`

## Install Gator
`go install github.com/sambakker4/gator@latest`

## Config
Create a file in your root directory called `.gatorconfig.json` and paste in the following contents
```
{
 "db_url": "postgres://postgres:@localhost:5432/gator?sslmode=disable",
 "current_user_name": ""
}
```

## Commands
`gator register <username>` registers and logins a user
`gator login <username>` logins specified user
`gator reset` resets database
`gator users` lists users
`gator agg <time_between_reqs>` aggregates feeds of the current user in an infinite loop
`gator addfeed <name> <url>` adds a feed by the current user
`gator feeds` lists all feeds
`gator follow <url>` makes the current user follow a feed
`gator following` lists all the feeds the current user is following
`gator unfollow <url>` unfollows the current user from specified feed
`gator browse <limit>(optional)` browses posts from feeds the current user is following
