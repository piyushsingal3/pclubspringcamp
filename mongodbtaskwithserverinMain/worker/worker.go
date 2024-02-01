// worker/worker.go

/*package worker

import (
	"log"
	"time"

	"mongodbtask/cfapi"
	"mongodbtask/store"
)

// PerformWork runs an infinite loop, opening connection with the database, getting data from Codeforces API,
// and storing it in the database. It then sleeps for 5 minutes.
func PerformWork(m *store.MongoStore) {
	for {
		// Open connection with the database
		err := m.OpenConnectionWithMongoDB("mongodb+srv://hello:hello123@cluster0.krf58nx.mongodb.net/?retryWrites=true&w=majority", "recentActions")
		if err != nil {
			log.Printf("Error opening connection with MongoDB: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}

		// Create a new Codeforces client
		codeforcesClient := cfapi.NewCodeforcesClient()

		// Get data from Codeforces API (implement your logic here)

		// Call the StoreRecentActionsInTheDatabase method
		// Assuming you have a function to fetch recent actions from Codeforces API
		recentActions, err := codeforcesClient.FetchRecentActions()
		if err != nil {
			log.Printf("Error fetching recent actions from Codeforces: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}

		// Store recent actions in the database
		err = m.StoreRecentActionsInTheDatabase(recentActions)
		if err != nil {
			log.Printf("Error storing recent actions in the database: %v", err)
		}

		// Close the connection with the database
		m.Close()

		// Sleep for 5 minutes before the next iteration
		time.Sleep(5 * time.Minute)
	}
}*/
// worker/worker.go

package worker

import (
	"log"
	"time"

	//"mongodbtask/models"
	"mongodbtask/cfapi"
	"mongodbtask/store"
)

func PerformWork(m *store.MongoStore) {
	for {

		if err := m.OpenConnectionWithMongoDB("mongodb://localhost:27017", "newdata"); err != nil {
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
