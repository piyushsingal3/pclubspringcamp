package cfapi

import (
	"encoding/json"
	"io"

	"fmt"
	"mongodbtask/models"
	"net/http"
)

// Create a client structure
type CodeforcesClient struct {
	client *http.Client
}

// Creates a new codeforces client
func NewCodeforcesClient() *CodeforcesClient {
	return &CodeforcesClient{
		client: http.DefaultClient,
	}
}

// fetch Data from codeforces api
func (c *CodeforcesClient) FetchRecentActions() (*models.Result, error) {

	response, err := http.Get("https://codeforces.com/api/recentActions?maxCount=30")
	if err != nil {
		return nil, fmt.Errorf("error fetching data from Codeforces API: %v", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var apiResponse models.Result
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return &apiResponse, nil
}

// func (c *CodeforcesClient) FetchUsersData() (*models.UserResult, error) {

// 	// response, err := http.Get("https://codeforces.com/api/user.ratedList?maxCount=1")
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("error fetching data from Codeforces API: %v", err)
// 	// }

// 	// defer response.Body.Close()

// 	// body, err := io.ReadAll(response.Body)
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("error reading response body: %v", err)
// 	// }

// 	// var userResponse models.UserResult
// 	// err = json.Unmarshal(body, &userResponse)
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
// 	// }

// 	// return &userResponse, nil
// }
