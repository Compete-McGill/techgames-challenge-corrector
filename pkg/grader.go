package corrector

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/Compete-McGill/techgames-challenge-corrector/pkg/util"

	"github.com/Compete-McGill/techgames-challenge-corrector/pkg/models"
)

var userInfo *models.CreateAccountRequest = &models.CreateAccountRequest{
	Email:    "example@email.com",
	Password: "password",
	FullName: "full name",
}

var articleInfo *models.CreateArticleRequest = &models.CreateArticleRequest{
	Title:    "test article",
	Subtitle: "test subtitle",
	Body:     "test body",
	UserID:   testUser.ID,
}

var testUser models.CreateAccountResponse
var testArticle models.CreateArticleResponse

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

	scores := make(map[string]bool)

	scores["liveness"] = livenessTest(userServer)
	scores["createAccount201"] = createAccount201Test(userServer)
	scores["createAccount400"] = createAccount400Test(userServer)
	scores["createAccount500"] = createAccount500Test(userServer)
	scores["createArticles200Test"] = createArticles200Test(userServer)
	scores["createArticles400Test"] = createArticles400Test(userServer)
	scores["indexArticlesTest"] = indexArticlesTest(userServer)
	scores["showArticles200Test"] = showArticles200Test(userServer)
	scores["showArticles404Test"] = showArticles404Test(userServer)
	scores["updateArticles200Test"] = updateArticles200Test(userServer)
	scores["updateArticles400Test"] = updateArticles400Test(userServer)
	scores["updateArticles404Test"] = updateArticles404Test(userServer)
	scores["deleteArticles200Test"] = deleteArticles200Test(userServer)
	scores["deleteArticles404Test"] = deleteArticles404Test(userServer)

	for test, score := range scores {
		status := ""
		if score {
			status = "passed"
		} else {
			status = "failed"
		}

		log.Printf("%s %s the %s test", userServer.name, status, test)
	}
}

func livenessTest(userServer *UserServer) bool {
	log.Println("Testing GET /api/status")
	resp, err := http.Get("http://localhost:" + userServer.port + "/api/status")
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	resp.Body.Close()

	expectedBody, _ := json.Marshal(&models.LivenessResponse{
		Status: "Up",
	})

	score, err := util.JSONBytesEqual(body, expectedBody)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	return score
}

func createAccount201Test(userServer *UserServer) bool {
	log.Println("Testing POST /api/auth/createAccount 201")

	userInfo.Email = userServer.name + "@email.com"
	userJSON, _ := json.Marshal(userInfo)

	resp, err := http.Post("http://localhost:"+userServer.port+"/api/auth/createAccount", "application/json", bytes.NewBuffer(userJSON))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	resp.Body.Close()

	err = json.Unmarshal(body, &testUser)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}

	return resp.Status == "201"
}

func createAccount400Test(userServer *UserServer) bool {
	log.Println("Testing POST /api/auth/createAccount 400")

	userJSON, _ := json.Marshal(&models.CreateAccountIncompleteRequest{
		Password: "password",
		FullName: "full name",
	})

	resp, err := http.Post("http://localhost:"+userServer.port+"/api/auth/createAccount", "application/json", bytes.NewBuffer(userJSON))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	resp.Body.Close()

	return resp.Status == "400"
}

func createAccount500Test(userServer *UserServer) bool {
	log.Println("Testing POST /api/auth/createAccount 500")

	userJSON, _ := json.Marshal(userInfo)

	resp, err := http.Post("http://localhost:"+userServer.port+"/api/auth/createAccount", "application/json", bytes.NewBuffer(userJSON))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	resp.Body.Close()

	return resp.Status == "500"
}

func createArticles200Test(userServer *UserServer) bool {
	log.Println("Testing POST /api/articles 200")

	articleJSON, _ := json.Marshal(articleInfo)

	resp, err := http.Post("http://localhost:"+userServer.port+"/api/articles", "application/json", bytes.NewBuffer(articleJSON))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	resp.Body.Close()

	err = json.Unmarshal(body, &testArticle)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}

	return resp.Status == "200"
}

func createArticles400Test(userServer *UserServer) bool {
	log.Println("Testing POST /api/articles 400")

	articleJSON, _ := json.Marshal(&models.CreateAccountIncompleteRequest{
		Password: "password",
		FullName: "full name",
	})

	resp, err := http.Post("http://localhost:"+userServer.port+"/api/articles", "application/json", bytes.NewBuffer(articleJSON))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	resp.Body.Close()

	return resp.Status == "400"
}

func indexArticlesTest(userServer *UserServer) bool {
	log.Println("Testing GET /api/articles 200")

	resp, err := http.Get("http://localhost:" + userServer.port + "/api/articles")
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	resp.Body.Close()

	var articles models.GetArticlesResponse
	err = json.Unmarshal(body, &articles)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}

	result := false
	for _, article := range articles {
		if article.UserID == testUser.ID && article.Title == testArticle.Title {
			result = true
		}
	}

	return result
}

func showArticles200Test(userServer *UserServer) bool {
	log.Println("Testing GET /api/articles/{articleId} 200")
	return false
}

func showArticles404Test(userServer *UserServer) bool {
	log.Println("Testing GET /api/articles/{articleId} 404")
	return false
}

func updateArticles200Test(userServer *UserServer) bool {
	log.Println("Testing PUT /api/articles/{atricleId} 200")
	return false
}

func updateArticles400Test(userServer *UserServer) bool {
	log.Println("Testing PUT /api/articles/{atricleId} 400")
	return false
}

func updateArticles404Test(userServer *UserServer) bool {
	log.Println("Testing PUT /api/articles/{atricleId} 404")
	return false
}

func deleteArticles200Test(userServer *UserServer) bool {
	log.Println("Testing DELETE /api/articles/{atricleId} 200")
	return false
}

func deleteArticles404Test(userServer *UserServer) bool {
	log.Println("Testing DELETE /api/articles/{atricleId} 404")
	return false
}
