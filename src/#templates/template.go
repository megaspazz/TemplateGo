package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func Init() {
}

func Solve(io *FastIO) {

}

type FastIO struct {
	reader *bufio.Reader
	writer *bufio.Writer
}

func (io *FastIO) NextChar() byte {
	c, _ := io.reader.ReadByte()
	return c
}

func (io *FastIO) NextInt() int {
	c := io.NextChar()
	for IsSpaceChar(c) {
		c = io.NextChar()
	}
	sgn := 1
	if c == '-' {
		sgn = -1
		c = io.NextChar()
	}
	res := 0
	for !IsSpaceChar(c) {
		res = (res * 10) + int(c-'0')
		c = io.NextChar()
	}
	return sgn * res
}

func (io *FastIO) NextLong() int64 {
	c := io.NextChar()
	for IsSpaceChar(c) {
		c = io.NextChar()
	}
	sgn := int64(1)
	if c == '-' {
		sgn = -1
		c = io.NextChar()
	}
	res := int64(0)
	for !IsSpaceChar(c) {
		res = (res * 10) + int64(c-'0')
		c = io.NextChar()
	}
	return sgn * res
}

func (io *FastIO) NextIntArray(size int) []int {
	return io.NextIntArrayOffset(size, 0)
}

func (io *FastIO) NextIntArrayOffset(size, offset int) []int {
	arr := make([]int, size+offset)
	for i := 0; i < size; i++ {
		arr[i+offset] = io.NextInt()
	}
	return arr
}

func (io *FastIO) NextLongArray(size int) []int64 {
	return io.NextLongArrayOffset(size, 0)
}

func (io *FastIO) NextLongArrayOffset(size, offset int) []int64 {
	arr := make([]int64, size+offset)
	for i := 0; i < size; i++ {
		arr[i+offset] = io.NextLong()
	}
	return arr
}

func (io *FastIO) NextString() string {
	c := io.NextChar()
	for IsSpaceChar(c) {
		c = io.NextChar()
	}
	var sb bytes.Buffer
	for !IsSpaceChar(c) {
		sb.WriteByte(c)
		c = io.NextChar()
	}
	return sb.String()
}

func (io *FastIO) NextLine() string {
	c := io.NextChar()
	for IsSpaceChar(c) {
		c = io.NextChar()
	}
	var sb bytes.Buffer
	for !IsLineBreakChar(c) {
		sb.WriteByte(c)
		c = io.NextChar()
	}
	return sb.String()
}

func (io *FastIO) Println(args ...interface{}) {
	io.writer.WriteString(fmt.Sprintln(args...))
}

func (io *FastIO) Printf(format string, args ...interface{}) {
	io.writer.WriteString(fmt.Sprintf(format, args...))
}

func (io *FastIO) FlushOutput() {
	io.writer.Flush()
}

func IsSpaceChar(c byte) bool {
	switch c {
	case 0, '\t', '\n', '\r', ' ':
		return true
	default:
		return false
	}
}

func IsLineBreakChar(c byte) bool {
	switch c {
	case 0, '\n', '\r':
		return true
	default:
		return false
	}
}

func main() {
	io := FastIO{reader: bufio.NewReader(os.Stdin), writer: bufio.NewWriter(os.Stdout)}
	Init()
	Solve(&io)
	io.FlushOutput()
}
