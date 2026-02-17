package test

import (
	"encoding/csv"
	"os"
	"testing"
)

// GetTestCsvReport retrieves records from test CSV report created during testing
func GetTestCsvReport(reportPath string, t *testing.T) [][]string {
	reportFile, err := os.Open(reportPath)

	if err != nil {
		t.Fatalf("Error occurred during reading test report file %v", err)
	}

	defer func() { _ = reportFile.Close() }()

	testReport, err := csv.NewReader(reportFile).ReadAll()

	if err != nil {
		t.Fatalf("Error occurred during reading test report records %v", err)
	}

	return testReport
}
