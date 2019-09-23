package io

import (
	"bufio"
	"io"
	"strings"
	"testing"
)

func TestMultiReader(t *testing.T) {
	r1 := strings.NewReader("one\n")
	r2 := strings.NewReader("two\n")
	reader := MultiReader(r1, r2)

	newReader := bufio.NewReader(reader)
ReadEnd:
	for {
		line, isPrefix, err := newReader.ReadLine()
		t.Log(string(line), isPrefix, err)
		if err == io.EOF {
			break ReadEnd
		}

	}
}
