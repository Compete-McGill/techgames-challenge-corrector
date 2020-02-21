package corrector

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// User contains information about a user
type User struct {
	Id             string  `json:"_id"`
	Email          string  `json:"email"`
	GithubToken    string  `json:"githubToken"`
	GithubUsername string  `json:"githubUsername"`
	GithubRepo     string  `json:"githubRepo"`
	Scores         []Score `json:"scores"`
}

type Score struct {
	Liveness          bool `json:"liveness"`
	IndexArticles     bool `json:"indexArticles"`
	ShowArticles200   bool `json:"showArticles200"`
	ShowArticles400   bool `json:"showArticles400"`
	ShowArticles404   bool `json:"showArticles404"`
	CreateArticles200 bool `json:"createArticles200"`
	CreateArticles400 bool `json:"createArticles400"`
	UpdateArticles200 bool `json:"updateArticles200"`
	UpdateArticles400 bool `json:"updateArticles400"`
	UpdateArticles404 bool `json:"updateArticles404"`
	DeleteArticles200 bool `json:"deleteArticles200"`
	DeleteArticles400 bool `json:"deleteArticles400"`
	DeleteArticles404 bool `json:"deleteArticles404"`
}

// UpdateScore ...
func UpdateScore(scores Score, userServer *UserServer) error {
	log.Printf("Getting user _id for user %s", userServer.name)
	response, err := http.Get(hostURL + "/username/" + userServer.name)
	if err != nil {
		log.Printf("Server HTTP request failed with error: %s\n", err)
		return err
	}
	if response.StatusCode != 200 {
		log.Printf("Bad GET request %d with error: %s\n", response.StatusCode, err)
		return err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Unable to read from response body: %s\n", err)
		return err
	}
	var data User
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("Error unmarshalling: %s\n", err)
		return err
	}

	log.Println("Updating user scores")
	b, err := json.Marshal(scores)
	if err != nil {
		log.Printf("Problem with Marshalling updateScore request: %s\n", err)
		return err
	}
	response, err = http.Post(hostURL+"/users/"+data.Id+"/updateScore", "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Printf("Server HTTP request failed with error: %s\n", err)
		return err
	}
	if response.StatusCode != 200 {
		log.Printf("Bad POST request %d with error: %s\n", response.StatusCode, err)
		return err
	}
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Unable to read from response body: %s\n", err)
		return err
	}
	log.Printf("UpdateScore request successfully sent")
	return nil
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
