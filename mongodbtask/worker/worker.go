package worker

import (
	"log"
	"mongodbtask/cfapi"
	"mongodbtask/store"
	"sync"
	"time"
)

func PerformWork(m *store.MongoStore, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// Open connection
		if err := m.OpenConnectionWithMongoDB("mongodb+srv://hello:hello123@cluster0.krf58nx.mongodb.net/?retryWrites=true&w=majority", "recent"); err != nil {
			log.Printf("Error opening connection with MongoDB: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}

		codeforcesClient := cfapi.NewCodeforcesClient()
		// Fetches data from api
		apiResponse, err := codeforcesClient.FetchRecentActions()
		if err != nil {
			log.Printf("Error fetching recent actions from Codeforces: %v", err)
			time.Sleep(10 * time.Second)
			m.Close()
			continue
		}
		// Stores data in mongodb database
		err = m.StoreRecentActionsInTheDatabase(apiResponse.RecentActions)
		if err != nil {
			log.Printf("Error storing recent actions in the database: %v", err)
		}

		m.Close()

		time.Sleep(5 * time.Minute)
	}
}
