package nesbconvertpin

import (
	"os"
	"path"
)

func CSMP_PIN_PARA() []byte {
	fileName := path.Join(getRootDir(), "etc/enum/CSMP_PIN_PARA.txt")
	buf, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return buf
}
