package decoder

import (
	"fmt"
	"l2/pkg/dictionary"
	"l2/pkg/reader"
	"l2/pkg/writer"

	unicode "l2/pkg/uniCode"
)

type Decoder struct {
	reader  *reader.Reader
	writer  *writer.Writer
	dict    *dictionary.Dictionary
	uniCode unicode.UniCode
}

func Decoder_createDecoder(reader *reader.Reader, writer *writer.Writer) *Decoder {
	decoder := &Decoder{
		reader: reader,
		writer: writer,
		dict:   dictionary.Dictionary_CreateDictionary()}

	return decoder

}

func (decoder *Decoder) decode() {

	size := decoder.uniCode.DecodeNumber()
	size--

	//let pk be the first code in the code stream
	k := decoder.uniCode.DecodeNumber()
	k--
	fmt.Println(k)
	//OUTPUY
	c := decoder.dict.Dictionary_GetKey(k)
	decoder.writer.Writer_addBytes(c)
	size -= len(c)

	prev_k := k
	var prev_c []byte
	iterations := 0
	// prev_c = append([]byte{}, c...)
	for size > 0 && decoder.reader.IsReading {
		k = decoder.uniCode.DecodeNumber()
		k--
		am := decoder.dict.Dictionary_GetKey(prev_k)
		prev_c = append([]byte{}, am...)

		// var S []byte
		if decoder.dict.Dictionary_IsValueContained(k) {
			// S = append(S, C)
			c = append([]byte{}, decoder.dict.Dictionary_GetKey(k)...)

			decoder.dict.Dictionary_AddKey(append(prev_c, c[0]))
			decoder.writer.Writer_addBytes(c)
			size -= len(c)
		} else {
			// S = decoder.dict.Dictionary_GetKey(k)
			decoder.dict.Dictionary_AddKey(append(prev_c, prev_c[0]))
			decoder.writer.Writer_addBytes(append(prev_c, prev_c[0]))
			size -= len(prev_c) + 1
		}
		// decoder.writer.Writer_addBytes(S)
		// size -= len(S)
		// C = S[0]
		// decoder.dict.Dictionary_AddKey(append(decoder.dict.Dictionary_GetKey(pk), C))
		prev_k = k
		iterations++
		if iterations == 1447 {
			fmt.Println("dupa")
		}
	}
	fmt.Println(iterations)

	decoder.writer.Writer_Flush()
}

func (decoder *Decoder) getCoding() {
	coding, err := decoder.reader.Reader_readByte()

	if err != nil {
		panic(err)
	}

	switch unicode.Coding(coding) {
	case unicode.Gamma:
		{
			decoder.uniCode = unicode.EliasGamma_Create(decoder.reader)
		}
	case unicode.Delta:
		{
			decoder.uniCode = unicode.EliasDelta_Create(decoder.reader)
		}
	case unicode.Omega:
		{
			decoder.uniCode = unicode.EliasOmega_Create(decoder.reader)
		}
	case unicode.Fib:
		{
			decoder.uniCode = unicode.Fibonacci_Create(decoder.reader)
		}
	}

}

func (decoder *Decoder) Decoder_run() {
	decoder.getCoding()
	decoder.decode()

}
