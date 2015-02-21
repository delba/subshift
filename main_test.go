package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestReplaceLines(t *testing.T) {
	t.Skip()

	srcPath := "test/fixtures/input.srt"
	dstPath := "test/fixtures/output.srt"
	expPath := "test/fixtures/expectation.srt"

	src, _ := os.Open(srcPath)
	defer src.Close()
	r := bufio.NewReader(src)

	dst, _ := os.Create(dstPath)
	defer dst.Close()
	defer os.Remove(dstPath)
	w := bufio.NewWriter(dst)

	ReplaceLines(r, w, time.Duration(5)*time.Second)

	output, _ := ioutil.ReadFile(dstPath)
	expectation, _ := ioutil.ReadFile(expPath)

	if string(output) != string(expectation) {
		t.Error("Nope")
	}
}
