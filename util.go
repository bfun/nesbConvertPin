package nesbconvertpin

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"os"
	"path"
)

func getRootDir() string {
	return os.Getenv("FAPWORKDIR")
}

func fileScanner(filename string) (file *os.File, scanner *bufio.Scanner) {
	var err error
	file, err = os.Open(path.Join(getRootDir(), filename))
	if err != nil {
		panic(err)
	}
	scanner = bufio.NewScanner(file)
	return
}
func getGbFileDecoder(fileName string) *xml.Decoder {
	buf, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	decoder := xml.NewDecoder(transform.NewReader(bytes.NewReader(buf), simplifiedchinese.GB18030.NewDecoder()))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		return transform.NewReader(input, simplifiedchinese.GB18030.NewDecoder()), nil
	}
	return decoder
}
