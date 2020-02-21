package main

import (
	corrector "github.com/Compete-McGill/techgames-challenge-corrector/pkg"
)

func main() {
	repos := []string{"https://github.com/Compete-McGill/techgames-sample-api.git", "https://github.com/Compete-McGill/techgames-sample-api.git", "https://github.com/Compete-McGill/techgames-sample-api.git", "https://github.com/Compete-McGill/techgames-sample-api.git", "https://github.com/Compete-McGill/techgames-sample-api.git"}

	// repos, err := corrector.GetTestInfo()
	// if err != nil {
	// 	log.Fatalf("Unable to get repos from challenge server: %s", err)
	// }

	users := corrector.Setup(repos)
	userServers := corrector.Run(users)
	corrector.Grade(userServers)
	corrector.Kill(userServers)
	corrector.Clean(users)
}
