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

var articleInfo *models.CreateArticleRequest = &models.CreateArticleRequest{
	Title:    "test article",
	Subtitle: "test subtitle",
	Body:     "test body",
	Author:   "test author",
}

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

	var testArticle models.CreateArticleResponse
	scores := make(map[string]bool)

	scores["liveness"] = livenessTest(userServer, &testArticle)
	scores["createArticles200Test"] = createArticles200Test(userServer, &testArticle)
	scores["createArticles400Test"] = createArticles400Test(userServer, &testArticle)
	scores["indexArticlesTest"] = indexArticlesTest(userServer, &testArticle)
	scores["showArticles200Test"] = showArticles200Test(userServer, &testArticle)
	scores["showArticles400Test"] = showArticles400Test(userServer, &testArticle)
	scores["showArticles404Test"] = showArticles404Test(userServer, &testArticle)
	scores["updateArticles200Test"] = updateArticles200Test(userServer, &testArticle)
	scores["updateArticles400Test"] = updateArticles400Test(userServer, &testArticle)
	scores["updateArticles404Test"] = updateArticles404Test(userServer, &testArticle)
	scores["deleteArticles200Test"] = deleteArticles200Test(userServer, &testArticle)
	scores["deleteArticles400Test"] = deleteArticles400Test(userServer, &testArticle)
	scores["deleteArticles404Test"] = deleteArticles404Test(userServer, &testArticle)

	total := 0
	for _, score := range scores {
		if score {
			total++
		}
	}
	log.Printf("%v's score: %v/13", userServer.name, total)
	// Make request to update user score
}

func livenessTest(userServer *UserServer, testArticle *models.CreateArticleResponse) bool {
	resp, err := http.Get("http://localhost:" + userServer.port + "/status")
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

func createArticles200Test(userServer *UserServer, testArticle *models.CreateArticleResponse) bool {
	articleJSON, _ := json.Marshal(articleInfo)

	resp, err := http.Post("http://localhost:"+userServer.port+"/articles", "application/json", bytes.NewBuffer(articleJSON))
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

	result := false
	if resp.StatusCode == 200 && testArticle.ID != "" && testArticle.Author == "test author" && testArticle.Title == articleInfo.Title {
		result = true
	}

	return result
}

func createArticles400Test(userServer *UserServer, testArticle *models.CreateArticleResponse) bool {
	articleJSON, _ := json.Marshal(&models.CreateArticleIncompleteRequest{
		Body:   "test body",
		Author: "test author",
	})

	resp, err := http.Post("http://localhost:"+userServer.port+"/articles", "application/json", bytes.NewBuffer(articleJSON))
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	resp.Body.Close()

	return resp.StatusCode == 400
}

func indexArticlesTest(userServer *UserServer, testArticle *models.CreateArticleResponse) bool {
	resp, err := http.Get("http://localhost:" + userServer.port + "/articles")
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
		if article.Author == "test author" && article.Title == testArticle.Title {
			result = true
		}
	}

	return result
}

func showArticles200Test(userServer *UserServer, testArticle *models.CreateArticleResponse) bool {
	resp, err := http.Get("http://localhost:" + userServer.port + "/articles/" + testArticle.ID)
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

	var article models.CreateArticleResponse
	err = json.Unmarshal(body, &article)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}

	return article.ID == testArticle.ID
}

func showArticles400Test(userServer *UserServer, testArticle *models.CreateArticleResponse) bool {
	resp, err := http.Get("http://localhost:" + userServer.port + "/articles/akshdkajhsd")
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	resp.Body.Close()

	return resp.StatusCode == 400
}

func showArticles404Test(userServer *UserServer, testArticle *models.CreateArticleResponse) bool {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:"+userServer.port+"/articles/507f1f77bcf86cd799439011", nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	resp.Body.Close()

	result := false
	if resp.StatusCode == 404 {
		result = true
	}

	return result
}

func updateArticles200Test(userServer *UserServer, testArticle *models.CreateArticleResponse) bool {
	updateArticleInfo := models.UpdateArticleRequest{Title: "new test title"}
	updateArticleJSON, _ := json.Marshal(&updateArticleInfo)

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, "http://localhost:"+userServer.port+"/articles/"+testArticle.ID, bytes.NewBuffer(updateArticleJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
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

	result := false
	if resp.StatusCode == 200 && testArticle.Title == updateArticleInfo.Title {
		result = true
	}

	return result
}

func updateArticles400Test(userServer *UserServer, testArticle *models.CreateArticleResponse) bool {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, "http://localhost:"+userServer.port+"/articles/hakjshdakjhdskj", nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	resp.Body.Close()

	result := false
	if resp.StatusCode == 400 {
		result = true
	}

	return result
}

func updateArticles404Test(userServer *UserServer, testArticle *models.CreateArticleResponse) bool {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, "http://localhost:"+userServer.port+"/articles/507f1f77bcf86cd799439011", nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	resp.Body.Close()

	result := false
	if resp.StatusCode == 404 {
		result = true
	}

	return result
}

func deleteArticles200Test(userServer *UserServer, testArticle *models.CreateArticleResponse) bool {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:"+userServer.port+"/articles/"+testArticle.ID, nil)
	deleteResp, err := client.Do(req)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	deleteResp.Body.Close()

	req, _ = http.NewRequest(http.MethodGet, "http://localhost:"+userServer.port+"/articles/"+testArticle.ID, nil)
	getResp, err := client.Do(req)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	getResp.Body.Close()

	result := false
	if getResp.StatusCode == 404 && deleteResp.StatusCode == 200 {
		result = true
	}

	return result
}

func deleteArticles400Test(userServer *UserServer, testArticle *models.CreateArticleResponse) bool {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:"+userServer.port+"/articles/akljshdakjshd", nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	resp.Body.Close()

	return resp.StatusCode == 400
}

func deleteArticles404Test(userServer *UserServer, testArticle *models.CreateArticleResponse) bool {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:"+userServer.port+"/articles/507f1f77bcf86cd799439011", nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return false
	}
	resp.Body.Close()

	result := false
	if resp.StatusCode == 404 {
		result = true
	}

	return result
}
