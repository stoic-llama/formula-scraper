package internal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/anaskhan96/soup"

	errHandler "github.com/stoic-llama/formula-scraper/pkg/errorHandlers"
)

func DoGet() string {
	// resp, err := soup.Get("https://www.target.com/s?searchTerm=enfamil+baby+formula")
	// if err != nil {
	// 	os.Exit(1)
	// }
	resp, err := http.Get("https://www.target.com/s?searchTerm=enfamil+baby+formula")
	
	errHandler.FatalError("Reading response body failed", err)
	body, _ := ioutil.ReadAll(resp.Body)
	doc := soup.HTMLParse(string(body))
	items := doc.FindAll("a", "class", "gCNFxQ") // doc.Find("div", "id", "comicLinks").FindAll("a")
	log.Println(len(items))

	return items[0].HTML()
	// resp, err := http.Get("https://www.target.com/s?searchTerm=enfamil+baby+formula")
	// errHandler.PrintError("Could not make get call", err)

	// body, err := ioutil.ReadAll(resp.Body)
	// errHandler.FatalError("Reading response body failed", err)

	// sb := string(body)
	// log.Printf("%s", sb)

	// return sb
}

func DoPost() {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Toby",
		"email": "Toby@example.com",
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)
	errHandler.FatalError("Failed on post call", err)
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	errHandler.FatalError("Failed on post call parse response body", err)
	sb := string(body)
	log.Printf("%s", sb)
}
