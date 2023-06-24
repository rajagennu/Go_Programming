package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	wordsBuffer := bytes.NewBufferString("word1 word2 word3 word4")
	expected := 4
	output := countWordsV2(wordsBuffer, false)
	if expected != output {
		t.Errorf("Expected %d, instead received %d", expected, output)
	}

}
