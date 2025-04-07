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
`go install github.com/sambakker4/gator`

## Set Up
### Enter the psql shell
### Mac `psql postgres`
### Linux `sudo -u postgres psql`

### Create a Database, run this in the psql shell
`CREATE DATABASE gator;`
### Change password (Linux Only)
`ALTER USER postgres PASSWORD 'postgres';`
