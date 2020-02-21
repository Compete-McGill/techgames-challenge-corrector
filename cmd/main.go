package main

import (
	"os"

	corrector "github.com/Compete-McGill/techgames-challenge-corrector/pkg"
)

func main() {

	apiURL := os.Getenv("API_URL")

	repos := []string{"https://github.com/Compete-McGill/techgames-sample-api.git", "https://github.com/Compete-McGill/techgames-sample-api.git", "https://github.com/Compete-McGill/techgames-sample-api.git", "https://github.com/Compete-McGill/techgames-sample-api.git", "https://github.com/Compete-McGill/techgames-sample-api.git"}

	// repos, err := corrector.GetTestInfo()
	// if err != nil {
	// 	log.Fatalf("Unable to get repos from challenge server: %s", err)
	// }

	users := corrector.Setup(repos, apiURL)
	userServers := corrector.Run(users)
	corrector.Grade(userServers)
	corrector.Kill(userServers)
	corrector.Clean(users)
}
