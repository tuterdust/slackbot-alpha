package main

import (
	"testing"
)

func TestGetENV(t *testing.T) {
	expectedResult := "/Users/tuterdust/go"
	actualResult := getenv("GOPATH")
	if expectedResult != actualResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}

}
