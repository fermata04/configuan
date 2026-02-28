package search

type SearchResult struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	Snippet string `json:"snippet"`
	Source  string `json:"source"`
}

func Search(query string) ([]SearchResult, error) {
	return []SearchResult{}, nil
}
