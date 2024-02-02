package main

import (
	"mongodbtask/server"
	"mongodbtask/store"
	"mongodbtask/worker"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	mongoStore := store.NewMongoStore()
	go worker.PerformWork(mongoStore, &wg)
	go server.Performserver(mongoStore, &wg)
	wg.Wait()

}
