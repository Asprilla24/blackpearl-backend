package model

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type SearchRequest struct {
	Pagination int64  `json:"pagination"`
	SearchWord string `json:"searchWord"`
}

type SearchResponse struct {
	Movies []Movie `json:"movies"`
	Err    string  `json:"err,omitempty"`
}

type SearchIMDBResponse struct {
	Search       []Movie `json:"Search"`
	TotalResults string  `json:"totalResults"`
	Response     string  `json:"Response"`
}
