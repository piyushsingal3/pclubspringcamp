package models

type BlogEntry struct {
	ID                      int    `json:"id"`
	Title                   string `json:"title"`
	OriginalLocal           string `json:"originalLocale"`
	CreationTimeSeconds     int    `json:"creationTimeSeconds"`
	Rating                  int    `json:"rating"`
	AuthorHandle            string `json:"authorHandle"`
	ModificationTimeSeconds int    `json:"modificationTimeSeconds"`
	Locale                  string `json:"locale"`
}

type RecentActions struct {
	TimeSeconds int       `json:"timeSeconds"`
	BlogEntry   BlogEntry `json:"blogEntry"`
	Comment     Comment   `json:"comment"`
	Rating      int       `json:"rating"`
}

type Comment struct {
	ID           int    `json:"id"`
	AuthorHandle string `json:"authorHandle"`
	Content      string `json:"content"`
	CreationTime int64  `json:"creationTime"`
	Rating       int    `json:"rating"`
	Locale       string `json:"locale"`
}

type Result struct {
	RecentActions []RecentActions `json:"result"`
}
