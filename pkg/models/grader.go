package models

// LivenessResponse is the format of a response from /status
type LivenessResponse struct {
	Status string `json:"status"`
}

// CreateAccountRequest is the format of the request to create a user successfully
type CreateAccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
}

// CreateAccountIncompleteRequest is the format of the request to fail with 400 when creating a user
type CreateAccountIncompleteRequest struct {
	Password string `json:"password"`
	FullName string `json:"fullName"`
}

// CreateAccountResponse is the response format when creating a new user
type CreateAccountResponse struct {
	ID       string `json:"_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
}

// AuthenticateRequest is the format of the authentication request body
type AuthenticateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateArticleRequest is the request format for creating a new article
type CreateArticleRequest struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Body     string `json:"body"`
	UserID   string `json:"userId"`
}

// CreateArticleIncompleteRequest is the request format to fail with a 400 when creating a new article
type CreateArticleIncompleteRequest struct {
	Body   string `json:"body"`
	UserID string `json:"userId"`
}

// CreateArticleResponse is the reponse format when fetching an article
type CreateArticleResponse struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Body     string `json:"body"`
	UserID   string `json:"userId"`
}

// GetArticlesResponse is the reponse format when fetching all articles
type GetArticlesResponse []CreateArticleResponse
