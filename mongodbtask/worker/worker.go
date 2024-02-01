package worker

import (
	"log"
	"time"

	"mongodbtask/cfapi"
	"mongodbtask/store"
)

func PerformWork(m *store.MongoStore) {
	for {

		if err := m.OpenConnectionWithMongoDB("mongodb+srv://hello:hello123@cluster0.krf58nx.mongodb.net/?retryWrites=true&w=majority", "recentActions"); err != nil {
			log.Printf("Error opening connection with MongoDB: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}

		codeforcesClient := cfapi.NewCodeforcesClient()

		apiResponse, err := codeforcesClient.FetchRecentActions()
		if err != nil {
			log.Printf("Error fetching recent actions from Codeforces: %v", err)
			time.Sleep(10 * time.Second)
			m.Close()
			continue
		}

		err = m.StoreRecentActionsInTheDatabase(apiResponse.RecentActions)
		if err != nil {
			log.Printf("Error storing recent actions in the database: %v", err)
		}

		m.Close()

		time.Sleep(5 * time.Minute)
	}
}
