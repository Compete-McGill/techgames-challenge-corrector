package main

import (
	corrector "github.com/Compete-McGill/techgames-challenge-corrector/pkg"
)

func main() {
	// Not ready for multiple repos yet (configure multiple port strategy)
	repos := []string{"https://github.com/devslopes-learn/simple-express-server.git"}

	users := corrector.Setup(repos)
	userServers := corrector.Run(users)
	corrector.Grade(userServers)
	corrector.Kill(userServers)
	corrector.Clean(users)
}
