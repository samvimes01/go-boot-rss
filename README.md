# Go rss aggregator 

## Setup environment

 - Install Docker
 - Install Go
 - create local .env file `cp .env.example .env`, optionally fill with own values

## First run on local machine

```sh
make initdb # creates docker network and volume and run a postgres container, and creates a database
make dbup # run migrations
```