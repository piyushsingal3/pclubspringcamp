package worker

import (
	"log"
	"mongodbtask/cfapi"
	"mongodbtask/store"
	"time"
)

func PerformWork(m *store.MongoStore) {

	for {
		// Open connection
		if err := m.OpenConnectionWithMongoDB("mongodb://localhost:27017", "project"); err != nil {
			log.Printf("Error opening connection with MongoDB: %v", err)

		}

		codeforcesClient := cfapi.NewCodeforcesClient()
		// Fetches data from api
		apiResponse, err := codeforcesClient.FetchRecentActions()
		if err != nil {
			log.Printf("Error fetching recent actions from Codeforces: %v", err)
			m.Close()

		}
		// Stores data in mongodb database
		err = m.StoreRecentActionsInTheDatabase(apiResponse.RecentActions)
		if err != nil {
			log.Printf("Error storing recent actions in the database: %v", err)
		}
		// userResponse, err := codeforcesClient.FetchUsersData()
		// if err != nil {
		// 	log.Printf("Error fetching user data from Codeforces: %v", err)
		// 	m.Close()

		// }
		// // Stores data in mongodb database
		// err = m.StoreUsersInTheDatabase(userResponse.Users)
		// if err != nil {
		// 	log.Printf("Error storing recent actions in the database: %v", err)
		// }

		time.Sleep(5 * time.Minute)
	}
}
