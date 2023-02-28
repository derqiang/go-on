package ch7

import (
	"bytes"
	"io"
)

type InterfaceValue int64

const debug = true

// 一个包含nil指针的接口，不是nil接口

func (InterfaceValue) Run() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
