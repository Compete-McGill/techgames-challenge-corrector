package main

import (
	"log"
	"os"

	corrector "github.com/Compete-McGill/techgames-challenge-corrector/pkg"
)

func main() {

	apiURL := os.Getenv("API_URL")
	secret := os.Getenv("SECRET")

	repos, err := corrector.GetTestInfo(apiURL, secret)
	log.Printf("Repos: %v", repos)
	if err != nil {
		log.Printf("Unable to get repos from challenge server: %s\n", err)
	}
	if len(repos) == 0 {
		log.Println("No repos, terminating script")
		return
	}
	// repos = []string{"https://github.com/Compete-McGill/techgames-sample-api.git", "https://github.com/Compete-McGill/techgames-sample-api.git", "https://github.com/Compete-McGill/techgames-sample-api.git", "https://github.com/Compete-McGill/techgames-sample-api.git", "https://github.com/Compete-McGill/techgames-sample-api.git"}
	// repos = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24"}

	for i := 0; i <= len(repos); i += 10 {
		rem := 10
		if len(repos) < i+10 {
			rem = len(repos) - i
		}
		currRepos := repos[i : rem+i]
		log.Printf("Running repos: %v", currRepos)

		users := corrector.Setup(currRepos, apiURL, secret)
		userServers := corrector.Run(users)
		corrector.Grade(userServers)
		corrector.Kill(userServers)
		corrector.Clean(users)
	}
}
