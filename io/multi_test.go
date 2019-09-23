package io

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestMultiReader(t *testing.T) {
	r1 := strings.NewReader("one\nthree\nfive\n")
	r2 := strings.NewReader("two\n")
	r3 := strings.NewReader("aaaaaaaaaaaaaaaa\ntwo\n")
	file, e := os.Open("D:\\video\\name.txt")
	if e != nil {
		return
	}
	reader := MultiReader(r1, r2, r3, file)

	newReader := bufio.NewReader(reader)
	for {
		line, isPrefix, err := newReader.ReadLine()
		log.Println("read output:", string(line), isPrefix, err)
		if err == io.EOF {
			break
		}

	}
}