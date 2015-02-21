package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/delba/subshift/converter"
)

var (
	timeline = regexp.MustCompile("-->")
	reader   *bufio.Reader
	writer   *bufio.Writer
	src      *os.File
	dst      *os.File
	srcPath  string
	dstPath  string
	delay    time.Duration
	err      error
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	srcPath = os.Args[1]

	delay = func(arg string) time.Duration {
		i, err := strconv.Atoi(arg)
		handle(err)
		return time.Duration(i)
	}(os.Args[2])

	dstPath = os.Args[3]

	// Create the reader

	src, err = os.Open(srcPath)
	handle(err)
	defer src.Close()

	reader = bufio.NewReader(src)

	// Create the writer

	dst, err = os.Create(dstPath)
	handle(err)
	defer dst.Close()

	writer = bufio.NewWriter(dst)

	// Execute the program

	ReplaceLines(reader, writer, delay)
}

func ReplaceLines(r *bufio.Reader, w *bufio.Writer, delay time.Duration) {
	for {
		line, err := r.ReadBytes('\n')

		if err == io.EOF && len(line) == 0 {
			break
		}

		if err != nil && err != io.EOF {
			fmt.Println(err)
			break
		}

		if timeline.Match(line) {
			line = converter.Format.ReplaceAllFunc(line, func(b []byte) []byte {
				d, err := converter.StringToDuration(string(b))
				handle(err)

				d += delay * time.Second
				s := converter.DurationToString(d)

				return []byte(s)
			})
		}

		w.Write(line)
	}

	w.Flush()
}
