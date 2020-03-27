package main

import (
	"github.com/gobuffalo/envy"
	"github.com/podded/maintainer"
)

func main() {

	envy.Load()

	mongoAddress := envy.Get("MONGO_ADDRESS", "localhost:27017")
	redisAddress := envy.Get("REDIS_ADDRESS", "localhost:6379")

	maint := maintainer.New(mongoAddress, redisAddress)

	maint.OrphanScrape()
}
