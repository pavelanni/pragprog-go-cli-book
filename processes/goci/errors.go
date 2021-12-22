package main

import (
	"errors"
	"fmt"
)

var (
	ErrValidation = errors.New("Validation failed")
)

type stepErr struct {
	step  string
	msg   string
	cause error
}

// Error returns a string to satisfy error interface
func (s *stepErr) Error() string {
	return fmt.Sprintf("Step: %q: %s: Cause: %v", s.step, s.msg, s.cause)
}

// Is compares errors
func (s *stepErr) Is(target error) bool {
	t, ok := target.(*stepErr)
	if !ok {
		return false
	}
	return t.step == s.step
}

// Unwrap shows the error details
func (s *stepErr) Unwrap() error {
	return s.cause
}
