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
	go func() {
		defer wg.Done()
		worker.PerformWork(mongoStore)
	}()
	go func() {
		defer wg.Done()
		server.Performserver(mongoStore)
	}()
	wg.Wait()
	mongoStore.Close()

}
