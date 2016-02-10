package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// http://marcio.io/2015/07/supercharging-atom-editor-for-go-development/

const rawurl = "https://qa-api.efset.org/test-results"
const key = "73c844c4-b851-424f-90b7-abad498f8463"

type apiResponse struct {
	TestResults []testResult `json:"test_results"`
	Next        string       `json:"next"`
}

type testResult struct {
	TestName   string `json:"test_name"`
	UserID     string `json:"user_id"`
	StartTime  string `json:"test_start_time"`
	FinishTime string `json:"test_finish_time"`
	Status     string `json:"status"`
	Scores     scores `json:"scores"`
}

type scores struct {
	Level    string `json:"level"`
	Combined string `json:"combined"`
}

func callAPI() {
	request := prepareRequest()
	client := &http.Client{}
	resp, err := client.Do(&request)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	parsed := parseResponse(body)
	writeCSV(parsed.TestResults)

	// fmt.Printf(string(resp.Header.Get("X-Api-Version")) + "\n")
}

func writeCSV(testResults []testResult) {
	headers := []string{
		"user_id",
		"test_name",
		"test_start_time",
		"test_finish_time",
		"status",
		"scores.combined",
		"scores.level",
	}

	w := csv.NewWriter(os.Stdout)
	w.Write(headers)

	for _, item := range testResults {
		records := []string{
			item.UserID,
			item.TestName,
			item.StartTime,
			item.FinishTime,
			item.Status,
			item.Scores.Combined,
			item.Scores.Level,
		}
		w.Write(records)
		// fmt.Printf("%v\n", item)
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

func prepareRequest() http.Request {

	u, err := url.Parse(rawurl)
	if err != nil {
		panic(err)
	}

	q := u.Query()
	q.Set("from", "2016-02-09T00:00:00.000Z")
	q.Set("to", "2116-01-01T00:00:00.000Z")
	u.RawQuery = q.Encode()
	fmt.Println("URL:>", u.String())

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("X-Api-Key", key)

	return *req
}

func parseResponse(body []byte) apiResponse {
	var m apiResponse
	err := json.Unmarshal(body, &m)

	if err != nil {
		panic(err)
	}

	return m
}

func main() {
	fmt.Printf("Hello, world.\n")
	callAPI()
}
