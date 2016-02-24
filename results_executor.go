package main

import (
	"io/ioutil"
	"net/http"
)

// CallParameters call parameters
type CallParameters struct {
	requestURL string
	secretKey  string
	fromDate   string
	toDate     string
}

var httpLib = new(HTTPLib)
var jsonLib = new(JSONLib)
var resultsWriter ResultsWriter = new(CSVResultsWriter)

//func Call(req http.Request, secretKey string) {

func Exec(params CallParameters) {
	req := httpLib.PrepareRequest(params.requestURL, params.fromDate, params.toDate, params.secretKey)
	execForRequest(req, params.secretKey)
}

func execForRequest(req http.Request, secretKey string) {

	resp := httpLib.MakeRequest(&req)

	if resp.StatusCode != 200 {
		panic(resp.Status)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	parsed := jsonLib.ParseResponse(body)

	if len(parsed.TestResults) > 0 {
		for _, item := range parsed.TestResults {
			record := []string{
				item.UserID,
				item.TestName,
				item.StartTime,
				item.FinishTime,
				item.Status,
				item.Scores.Combined,
				item.Scores.Level,
			}
			resultsWriter.Write(record)
		}

		if len(parsed.Next) > 0 {
			next := httpLib.PrepareRequestWhenFullURL(parsed.Next, secretKey)
			execForRequest(next, secretKey)
		}
	}
}
