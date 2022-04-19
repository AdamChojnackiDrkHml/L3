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
	var input string
	var output string
	var coding unicode.Coding
	var readerA *reader.Reader
	var writerA *writer.Writer
	fmt.Println(os.Args)
	if len(os.Args) < 3 {
		input = "data/input/test"
		output = "data/output/test"
		coding = unicode.Gamma
	} else {
		if len(os.Args) < 4 {
			os.Args = append(os.Args, "0")
		}
		input = os.Args[1]
		output = os.Args[2]

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

	readerA = reader.Reader_createReader(input)
	writerA = writer.Writer_createWriter(output)

	coder := coder.Coder_createCoder(readerA, writerA, coding)

	coder.Coder_run()

	writerA.CloseFile()

	f1, _ := os.Stat(os.Args[1])
	f2, _ := os.Stat(os.Args[2])
	fmt.Println(input, output, coding)
	fmt.Println("Rozmiar Wejścia:", f1.Size())
	fmt.Println("Rozmiar Wyjścia:", f2.Size())
	fmt.Println("Stopień kompresji: ", float64(100.0*(1.0-float64(f2.Size())/float64(f1.Size()))))
	coder.Coder_scanFile(os.Args[1])
	coder.Coder_scanFile(os.Args[2])

}
