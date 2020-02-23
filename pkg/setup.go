package corrector

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

var hostURL string
var secret string

// Setup clones repos and installs their dependencies
func Setup(repos []string, apiURL string, secr string) []string {
	hostURL = strings.Trim(apiURL, "/")
	secret = secr

	users := make([]string, 0, len(repos))
	var wg sync.WaitGroup
	for _, repo := range repos {
		wg.Add(1)
		go setupHelper(repo, &users, &wg)
	}

	wg.Wait()

	return users
}

func setupHelper(repo string, users *[]string, wg *sync.WaitGroup) {
	defer wg.Done()

	user := strings.Split(repo, "/")[3]

	// TODO: Add error handling
	log.Printf("Cloning %v\n", repo)
	exec.Command("git", "clone", repo, os.Getenv("HOME")+"/test-repos/"+user).Run()
	log.Printf("Installing dependencies for %v's server\n", user)
	exec.Command("npm", "install", os.Getenv("HOME")+"/test-repos/"+user, "--prefix", os.Getenv("HOME")+"/test-repos/"+user).Run()
	*users = append(*users, user)
}

// Clean removes the user repositories
func Clean(users []string) {
	log.Printf("Cleaning test directory\n")
	exec.Command("rm", "-rf", os.Getenv("HOME")+"/test-repos").Run()
	log.Printf("Done\n")
}
