package main

import (
	"encoding/csv"
	"log"
	"os"
)

var headers = []string{
	"user_id",
	"test_name",
	"test_start_time",
	"test_finish_time",
	"status",
	"scores.combined",
	"scores.level",
}

// WriteHeaders writes headers to Stdout
func WriteHeaders() {
	writer := csv.NewWriter(os.Stdout)

	writer.Write(headers)
	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

// WriteCSV writes CSV to Stdout
func WriteCSV(testResults []TestResult) {

	writer := csv.NewWriter(os.Stdout)

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
		writer.Write(records)
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}
