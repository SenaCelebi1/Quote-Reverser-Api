package model

type Quotes struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

type QuotesNew struct {
	Author string   `json:"author"`
	Quotes []string `json:"quotes"`
}
