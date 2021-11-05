package main

import "testing"

func TestNewCar(t *testing.T) {
	car := NewCar()
	if car == nil {
		t.Fail()
	}
}
