package main

import "testing"

func TestPop(t *testing.T) {
	// Arrange
	args := []string{"one", "two", "three"}
	var s1, s2, s3, s4 string

	// Act
	s1, args = pop(args)
	s2, args = pop(args)
	s3, args = pop(args)
	s4, args = pop(args)

	// Assert
	if s1 != "one" {
		t.Errorf("unexpected \"%s\"", s1)
	}
	if s2 != "two" {
		t.Errorf("unexpected \"%s\"", s2)
	}
	if s3 != "three" {
		t.Errorf("unexpected \"%s\"", s3)
	}
	if s4 != "" {
		t.Errorf("unexpected \"%s\"", s4)
	}
}