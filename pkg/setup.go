package corrector

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// Setup clones repos and installs their dependencies
func Setup(repos []string) []string {
	users := make([]string, 0, len(repos))
	// TODO: Upgrade to goroutine
	for _, repo := range repos {
		user := strings.Split(repo, "/")[3]
		users = append(users, user)

		log.Printf("Cloning %v\n", repo)
		exec.Command("git", "clone", repo, os.Getenv("HOME")+"/test-repos/"+user).Run()
		log.Printf("Installing dependencies for %v's server\n", user)
		exec.Command("npm", "install", "--prefix", os.Getenv("HOME")+"/test-repos/"+user).Run()
	}

	return users
}

// Clean removes the user repositories
func Clean(users []string) {
	log.Printf("Cleaning test directory\n")
	exec.Command("rm", "-rf", os.Getenv("HOME")+"/test-repos").Run()
	log.Printf("Done\n")
}
