package models

// LivenessResponse is the format of a response from /api/status
type LivenessResponse struct {
	Status string `json:"status"`
}

// CreateAccountCompleteRequest is the format of the request to create a user successfully
type CreateAccountCompleteRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
}

// CreateAccountIncompleteRequest  is the format of the request to fail with 400 when creating a user
type CreateAccountIncompleteRequest struct {
	Password string `json:"password"`
	FullName string `json:"fullName"`
}

type CreateAccountResponse struct {
	UserId   string `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
}

// AuthenticateRequest is the format of the authentication request body
type AuthenticateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateArticleCompleteRequest struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Body     string `json:"body"`
	UserId   string `json:"userId"`
}

type CreateArticleIncompleteRequest struct {
	Body   string `json:"body"`
	UserId string `json:"userId"`
}

type CreateArticleResponse struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Body     string `json:"body"`
	UserId   string `json:"userId"`
}
