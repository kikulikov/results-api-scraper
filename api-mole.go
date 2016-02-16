package main

import (
	"flag"
	"io/ioutil"
	"net/http"
)

// http://marcio.io/2015/07/supercharging-atom-editor-for-go-development/

const rawurl = "https://qa-api.efset.org/test-results"

var secretKey string
var fromDate string
var toDate string

func callAPI(req http.Request) {
	resp := MakeRequest(&req)

	if resp.StatusCode != 200 {
		panic(resp.Status)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	parsed := ParseResponse(body)

	if len(parsed.TestResults) > 0 {
		WriteCSV(parsed.TestResults)

		if len(parsed.Next) > 0 {
			// fmt.Println(PrepareRequestWhenFullURL(parsed.Next, secretKey))
			callAPI(PrepareRequestWhenFullURL(parsed.Next, secretKey))
		}
	}
}

func main() {
	flag.StringVar(&secretKey, "key", "", "API Key")
	flag.StringVar(&fromDate, "from", "", "Date FROM")
	flag.StringVar(&toDate, "to", "", "Date TO")
	flag.Parse()

	WriteHeaders()
	callAPI(PrepareRequest(rawurl, fromDate, toDate, secretKey))
}
