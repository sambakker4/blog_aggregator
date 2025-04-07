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
