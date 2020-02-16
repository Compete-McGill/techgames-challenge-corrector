package models

// LivenessResponse is the format of a response from /status
type LivenessResponse struct {
	Status string `json:"status"`
}

// CreateArticleRequest is the request format for creating a new article
type CreateArticleRequest struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Body     string `json:"body"`
	Author   string `json:"author"`
}

// CreateArticleIncompleteRequest is the request format to fail with a 400 when creating a new article
type CreateArticleIncompleteRequest struct {
	Body   string `json:"body"`
	Author string `json:"author"`
}

// CreateArticleResponse is the reponse format when fetching an article
type CreateArticleResponse struct {
	ID       string `json:"_id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Body     string `json:"body"`
	Author   string `json:"author"`
}

// GetArticlesResponse is the reponse format when fetching all articles
type GetArticlesResponse []CreateArticleResponse

// UpdateArticleRequest is the request format for updating an article
type UpdateArticleRequest struct {
	Title string `json:"title"`
}
