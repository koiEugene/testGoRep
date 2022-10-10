package main

import "testing"

func TestAdd(t *testing.T) {
	var x, y, result = 2, 2, 4
	realResult := abc(x, y)
	if realResult != realResult {
		t.Errorf("%d != %d", realResult, result)
	}

}
