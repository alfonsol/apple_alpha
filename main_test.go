package main

import "testing"

func TestDummy(t *testing.T) {
	if 1 != 2 {
		t.Errorf("Dummy test failed, should have passed")
	}
}
