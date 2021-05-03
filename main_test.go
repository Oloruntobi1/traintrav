package main

import "testing"

func TestSum(t *testing.T){
	
	got := Sum(7, 9)
	want := 16

	if got != want {
		t.Errorf("got %d, not equal to want %d", got, want)
	}

}