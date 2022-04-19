package decoder

import (
	"l2/pkg/reader"
	"l2/pkg/writer"
	"testing"
)

func TestDecode(t *testing.T) {
	r := reader.Reader_createReader("/home/adam/Dokumenty/Github/L3/data/output/test")
	w := writer.Writer_createWriter("/home/adam/Dokumenty/Github/L3/data/outputOutput/test")

	d := Decoder_createDecoder(r, w)

	d.Decoder_run()

}
