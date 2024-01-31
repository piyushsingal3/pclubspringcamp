package main

import (
	"mongodbtask/store"
	"mongodbtask/worker"
	"time"
)

func main() {

	mongoStore := store.NewMongoStore()

	go worker.PerformWork(mongoStore)

	select {
	case <-time.After(time.Hour):

	}
}
