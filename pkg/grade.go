package corrector

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Grade grades the user's server based on a series of tests
func Grade(userServers []*UserServer) error {
	// TODO: Upgrade to goroutine
	for _, userServer := range userServers {
		log.Println("Testing endpoint /")
		resp, err := http.Get("http://localhost:" + userServer.port)
		if err != nil {
			return err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		resp.Body.Close()

		log.Printf("Response: %+v\n", string(body))
	}

	return nil
}
