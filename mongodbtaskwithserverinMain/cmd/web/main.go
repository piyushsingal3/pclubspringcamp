package main

import (
	"fmt"
	"mongodbtask/store"
	"mongodbtask/worker"
	"time"
)

func main() {

	mongoStore := store.NewMongoStore()

	go worker.PerformWork(mongoStore)

	select {
	case <-time.After(10 * time.Minute):

		fmt.Println("10 minutes completed")

	}
}
