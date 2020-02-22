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
	repos = []string{"https://github.com/Compete-McGill/techgames-sample-api.git", "https://github.com/Compete-McGill/techgames-sample-api.git", "https://github.com/Compete-McGill/techgames-sample-api.git", "https://github.com/Compete-McGill/techgames-sample-api.git", "https://github.com/Compete-McGill/techgames-sample-api.git"}

	users := corrector.Setup(repos, apiURL, secret)
	userServers := corrector.Run(users)
	corrector.Grade(userServers)
	corrector.Kill(userServers)
	corrector.Clean(users)
}
