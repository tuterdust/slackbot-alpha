package main_test

import "testing"

func TestOnePlusOne(t *testing.T) {
	if 1+1 != 2 {
		t.Fatalf("Da faq")
	}
}
