package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 words2 word3 word4\n")
	exp := 4

	res := count(b, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2\nword3 word4\nline3\nline4")
	exp := 4
	res := count(b, true)

	if exp != res {
		t.Errorf("Expected %d, got %d instead\n", exp, res)
	}
}
