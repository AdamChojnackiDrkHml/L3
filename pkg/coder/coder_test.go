package coder

import (
	"l2/pkg/reader"
	unicode "l2/pkg/uniCode"
	"l2/pkg/writer"
	"testing"
)

func TestCode(t *testing.T) {
	r := reader.Reader_createReader("/home/adam/Dokumenty/Github/L3/data/input/SamplePDFFile_5mb.pdf")
	w := writer.Writer_createWriter("/home/adam/Dokumenty/Github/L3/data/output/test")

	c := Coder_createCoder(r, w, unicode.Gamma)

	c.Coder_run()

}
