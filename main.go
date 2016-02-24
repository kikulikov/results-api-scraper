package main

import "flag"

// http://marcio.io/2015/07/supercharging-atom-editor-for-go-development/

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

	Exec(CallParameters{requestURL, secretKey, fromDate, toDate})
}
