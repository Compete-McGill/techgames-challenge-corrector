package corrector

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// Grade grades the user's server based on a series of tests
func Grade(userServers []*UserServer) {
	var wg sync.WaitGroup

	for _, userServer := range userServers {
		wg.Add(1)
		go gradeHelper(userServer, &wg)
	}

	wg.Wait()
}

func gradeHelper(userServer *UserServer, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println("Testing endpoint /")
	resp, err := http.Get("http://localhost:" + userServer.port)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	resp.Body.Close()

	log.Printf("Response: %+v\n", string(body))
}
