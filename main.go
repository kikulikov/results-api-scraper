package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"strconv"
)

// http://marcio.io/2015/07/supercharging-atom-editor-for-go-development/

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

func main() {
	var requestURL string
	flag.StringVar(&requestURL, "url", "", "URL")

	var secretKey string
	flag.StringVar(&secretKey, "key", "", "API Key")

	var fromDate string
	flag.StringVar(&fromDate, "from", "", "Date FROM")

	var toDate string
	flag.StringVar(&toDate, "to", "", "Date TO")

	flag.Parse()

	// var httpLib = new(HTTPLib)
	// var jsonLib = new(JSONLib)
	// var resultsWriter = new(CSVResultsWriter)

	callOnParams(CallParameters{requestURL, secretKey, fromDate, toDate})
}

func callOnParams(params CallParameters) {
	req := httpLib.PrepareRequest(params.requestURL, params.fromDate, params.toDate, params.secretKey)
	callOnRequest(req, params.secretKey)
}

func callOnRequest(req http.Request, secretKey string) {

	body := retrieveBody(req)

	parsed := jsonLib.ParseResponse(body)

	if len(parsed.TestResults) > 0 {
		for _, item := range parsed.TestResults {
			record := []string{
				item.UserID,
				item.TestName,
				item.StartTime,
				item.FinishTime,
				item.Status,
				item.Scores.Level,
				strconv.Itoa(item.Scores.Score),
				strconv.Itoa(item.Scores.Reading.Score),
				strconv.Itoa(item.Scores.Listening.Score),
			}
			resultsWriter.Write(record)
		}

		if len(parsed.Next) > 0 {
			next := httpLib.PrepareRequestWhenFullURL(parsed.Next, secretKey)
			callOnRequest(next, secretKey)
		}
	}
}

func retrieveBody(req http.Request) []byte {
	resp := httpLib.MakeRequest(&req)

	if resp.StatusCode != 200 {
		panic(resp.Status)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return body
}
