package main

import (
	"net/http"
	"net/url"
)

// PrepareRequestWhenFullURL builds a request when full URL provided
func PrepareRequestWhenFullURL(rawurl string, key string) http.Request {

	u, err := url.Parse(rawurl)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("X-Api-Key", key)

	return *req
}

// PrepareRequest builds a request
func PrepareRequest(rawurl string, from string, to string, key string) http.Request {

	u, err := url.Parse(rawurl)
	if err != nil {
		panic(err)
	}

	q := u.Query()
	q.Set("from", from)
	q.Set("to", to)
	u.RawQuery = q.Encode()
	// fmt.Println("URL:>", u.String())

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("X-Api-Key", key)

	return *req
}

// MakeRequest makes a request. Doesn't close the body.
func MakeRequest(req *http.Request) *http.Response {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	return resp
}
