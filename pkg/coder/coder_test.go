package coder

import (
	"l2/pkg/reader"
	unicode "l2/pkg/uniCode"
	"l2/pkg/writer"
	"testing"
)

func TestCode(t *testing.T) {
	r := reader.Reader_createReader("/home/adam/Dokumenty/Github/L3/data/input/testy1/pan-tadeusz-czyli-ostatni-zajazd-na-litwie.txt")
	w := writer.Writer_createWriter("/home/adam/Dokumenty/Github/L3/data/output/test")

	c := Coder_createCoder(r, w, unicode.Fib)

	c.Coder_run()

}
