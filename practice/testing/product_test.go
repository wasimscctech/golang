package main

// notes: the file name must end with _test

import "testing"

// func TestProductofNothing(t *testing.T) { //function name must starts with "Test"
// 	if product() != 1 {
// 		t.Errorf("expected to be 1")
// 	}
// }

// func TestProductofaNumber(t *testing.T) {
// 	if product(15) != 15 {
// 		t.Errorf("expected to be 15")
// 	}
// }

func TestProductofTwoNumber(t *testing.T) {
	if product(15, 2) != 30 {
		t.Errorf("expected to be 30")
	}
}
