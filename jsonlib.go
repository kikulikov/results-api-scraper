package main

import "encoding/json"

// APIResponse APIResponse
type APIResponse struct {
	TestResults []TestResult `json:"test_results"`
	Next        string       `json:"next"`
}

// TestResult TestResult
type TestResult struct {
	TestName   string `json:"test_name"`
	UserID     string `json:"user_id"`
	StartTime  string `json:"test_start_time"`
	FinishTime string `json:"test_finish_time"`
	Status     string `json:"status"`
	Scores     Scores `json:"scores"`
}

// Scores Scores
type Scores struct {
	Level    string `json:"level"`
	Combined string `json:"combined"`
}

// ParseResponse unmarshals the json
func ParseResponse(body []byte) APIResponse {
	var m APIResponse
	err := json.Unmarshal(body, &m)

	if err != nil {
		panic(err)
	}

	return m
}
