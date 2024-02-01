package cfapi

import (
	"encoding/json"
	"io/ioutil"

	"fmt"
	"mongodbtask/models"
	"net/http"
)

type CodeforcesClient struct {
	client *http.Client
}

func NewCodeforcesClient() *CodeforcesClient {
	return &CodeforcesClient{
		client: http.DefaultClient,
	}
}

func (c *CodeforcesClient) FetchRecentActions() (*models.Result, error) {

	response, err := http.Get("https://codeforces.com/api/recentActions?maxCount=30")
	if err != nil {
		return nil, fmt.Errorf("error fetching data from Codeforces API: %v", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
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
