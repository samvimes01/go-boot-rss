# Go rss aggregator 

App that allows users subscribing to multiple rss2 feeds.
Auth is simplifid, done with ApiKey header
No 3d party routers used, since Go 1.22 intruduced new HandleFunc functionality.
(Only 3d party is viper lib for getting config)

## Setup environment

 - Install Docker (required only for local Postgresql)
 - Install Go
 - create local .env file `cp .env.example .env`, optionally fill with own values

## First run on local machine

```sh
make initdb # creates docker network and volume and run a postgres container, and creates a database
make dbup # run migrations
```

## Next run on local db

1. `make startdb` to start postgres docker container if not running
2. `make build`
3. `./server` to run the server
