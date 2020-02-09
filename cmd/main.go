package main

import (
	"log"

	corrector "github.com/Compete-McGill/techgames-challenge-corrector/pkg"
)

func main() {
	repos := []string{"https://github.com/rwieruch/node-express-server.git", "https://github.com/MohamedBeydoun/node-express-server.git"}

	users := corrector.Setup(repos)
	userServers := corrector.Run(users)

	if err := corrector.Grade(userServers); err != nil {
		log.Fatal(err)
	}
	if err := corrector.Kill(userServers); err != nil {
		log.Fatal(err)
	}

	corrector.Clean(users)
}
