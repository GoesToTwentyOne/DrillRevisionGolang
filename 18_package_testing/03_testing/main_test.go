package main

import "testing"

func TestIntMin(t *testing.T) {
	if IntMin(2, -2) != -2 {
		t.Error("fail")

	}

}
func TestIntMinTableDriven(t *testing.T) {
	if IntMinTableDriven() != true {
		t.Error("not ewual")

	}

}
