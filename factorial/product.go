package main

import "testing"

func test_productofnothing(t *testing.T) {
	if product() != 1 {
		t.Errorf("expected product is 1")
	}
}
