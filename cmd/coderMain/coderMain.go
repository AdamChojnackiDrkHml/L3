package main

import (
	"fmt"
	"l2/pkg/coder"
	"l2/pkg/reader"
	unicode "l2/pkg/uniCode"
	"l2/pkg/writer"
	"os"
	"strconv"
)

func main() {
	// fmt.Println(os.Getwd())
	var readerA *reader.Reader
	var writerA *writer.Writer
	var coding unicode.Coding
	fmt.Println(os.Args)
	if len(os.Args) != 4 {
		readerA = reader.Reader_createReader("data/input/test")
		writerA = writer.Writer_createWriter("data/output/test")
		coding = unicode.Gamma
	} else {
		readerA = reader.Reader_createReader(os.Args[1])
		writerA = writer.Writer_createWriter(os.Args[2])
		codingNo, err := strconv.Atoi(os.Args[3])

		if err != nil {
			codingNo = 0
		}
		switch unicode.Coding(codingNo) {

		case unicode.Gamma:
			{
				coding = unicode.Gamma
			}
		case unicode.Delta:
			{
				coding = unicode.Delta
			}
		case unicode.Omega:
			{
				coding = unicode.Omega
			}
		case unicode.Fib:
			{
				coding = unicode.Fib
			}

		}

	}

	coder := coder.Coder_createCoder(readerA, writerA, coding)

	coder.Coder_run()

	writerA.CloseFile()

	f1, _ := os.Stat(os.Args[1])
	f2, _ := os.Stat(os.Args[2])
	fmt.Println("Rozmiar Wejścia:", f1.Size())
	fmt.Println("Rozmiar Wyjścia:", f2.Size())
	fmt.Println("Stopień kompresji: ", float64(100.0*(1.0-float64(f2.Size())/float64(f1.Size()))))
	coder.Coder_scanFile(os.Args[1])
	coder.Coder_scanFile(os.Args[2])

	// coder.Coder_avgCodingLenght()
}
