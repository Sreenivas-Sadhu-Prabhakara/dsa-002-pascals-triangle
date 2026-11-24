package main

import (
	"reflect"
	"testing"
)

func TestPascal(t *testing.T) {
	got := Pascal(5)
	want := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Pascal(5) = %v, want %v", got, want)
	}
	if len(Pascal(0)) != 0 {
		t.Errorf("Pascal(0) should be empty")
	}
	if !reflect.DeepEqual(Pascal(1), [][]int{{1}}) {
		t.Errorf("Pascal(1) wrong")
	}
}
