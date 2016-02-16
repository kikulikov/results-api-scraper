package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"time"
)

// http://marcio.io/2015/07/supercharging-atom-editor-for-go-development/

const rawurl = "https://qa-api.efset.org/test-results"
const dateFormat = "2006-01-02T15:04:05.000Z07:00"

var secretKey string
var fromDate string
var toDate string
var tailInterval int

func main() {
	flag.StringVar(&secretKey, "key", "", "API Key")
	flag.StringVar(&fromDate, "from", "", "Date FROM")
	flag.StringVar(&toDate, "to", "", "Date TO")
	flag.IntVar(&tailInterval, "tail", 0, "Tail Interval (s)")
	flag.Parse()

	WriteHeaders()

	if tailInterval > 0 {
		var now = time.Now().UTC()
		var begin = now.Add(time.Duration(-1*tailInterval) * time.Second).Format(dateFormat)
		var end = now.Format(dateFormat)

		for {
			callAPI(PrepareRequest(rawurl, begin, end, secretKey))
			time.Sleep(time.Duration(tailInterval) * time.Second)
		}
	} else {
		callAPI(PrepareRequest(rawurl, fromDate, toDate, secretKey))
	}
}

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
