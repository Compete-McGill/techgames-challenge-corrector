package corrector

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// User contains information about a user
type User struct {
	Email          string  `json:"email"`
	GithubToken    string  `json:"githubToken"`
	GithubUsername string  `json:"githubUsername"`
	GithubRepo     string  `json:"githubRepo"`
	Scores         []Score `json:"scores"`
}

type Score struct {
	Liveness          string `json:"liveness"`
	Authenticate200   bool   `json:"authenticate200"`
	Authenticate403   bool   `json:"authenticate403"`
	CreateAccount201  bool   `json:"createAccount201"`
	CreateAccount400  bool   `json:"createAccount400"`
	CreateAccount500  bool   `json:"createAccount500"`
	IndexArticles     bool   `json:"indexArticles"`
	ShowArticles200   bool   `json:"showArticles200"`
	ShowArticles404   bool   `json:"showArticles404"`
	CreateArticles201 bool   `json:"createArticles201"`
	CreateArticles400 bool   `json:"createArticles400"`
	CreateArticles403 bool   `json:"createArticles403"`
	UpdateArticles200 bool   `json:"updateArticles200"`
	UpdateArticles400 bool   `json:"updateArticles400"`
	UpdateArticles401 bool   `json:"updateArticles401"`
	UpdateArticles403 bool   `json:"updateArticles403"`
	UpdateArticles404 bool   `json:"updateArticles404"`
	Timestamp         string `json:"timestamp"`
	DeleteArticles200 bool   `json:"deleteArticles200"`
	DeleteArticles401 bool   `json:"deleteArticles401"`
	DeleteArticles403 bool   `json:"deleteArticles403"`
	DeleteArticles404 bool   `json:"deleteArticles404"`
}

// GetTestInfo returns the URLs of the Github repos from the challenge API
func GetTestInfo() (repos []string, err error) {
	log.Println("Fetching user server data")
	response, err := http.Get("http://localhost:3000/users")
	if err != nil {
		log.Printf("Server HTTP request failed with error: %s\n", err)
		return nil, err
	}
	if response.StatusCode != 200 {
		log.Printf("Bad request %d with error: %s\n", response.StatusCode, err)
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Unable to read from response body: %s\n", err)
		return nil, err
	}
	var data []*User
	log.Println("Response body: " + string(body))
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("Error unmarshalling: %s\n", err)
		return nil, err
	}
	repos = []string{}
	for _, user := range data {
		repos = append(repos, user.GithubRepo)
	}
	log.Printf("Returned repos: %+v\n", repos)
	return repos, nil
}
