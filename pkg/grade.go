package corrector

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Grade grades the user's server based on a series of tests
func Grade(userServers []*UserServer) {
	// TODO: Upgrade to goroutine
	for _, userServer := range userServers {
		fmt.Println("Testing endpoint /ingredients")
		resp, err := http.Get("http://localhost:" + strconv.Itoa(userServer.port) + "/ingredients")
		if err != nil {
			fmt.Print(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Print(err)
		}
		resp.Body.Close()

		fmt.Printf("Response: %+v\n", string(body))
	}
}
