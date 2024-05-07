package nesbconvertpin

import (
	"bufio"
	"bytes"
	"os"
	"path"
	"strings"
)

type PinElem struct {
	Pin string
	Acc string
}

func CSMP_PIN_ELEM() (elems map[string][]PinElem) {
	elems = make(map[string][]PinElem)
	para := CSMP_PIN_PARA()
	fileName := path.Join(getRootDir(), "etc/enum/CSMP_PIN_ELEM.txt")
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "enumData1") || strings.HasPrefix(line, "NESB.") || !strings.Contains(line, ".NESB") {
			continue
		}
		kv := strings.Split(line, "\t")
		if len(kv) != 2 {
			continue
		}
		if !bytes.Contains(para, []byte("\n"+kv[0]+"\t")) {
			panic(line)
		}
		k := strings.Replace(kv[0], ".NESB", "", 1)
		var pes []PinElem
		couples := strings.Split(kv[1], "|")
		for _, couple := range couples {
			pa := strings.Split(couple, ",")
			if len(pa) < 2 {
				panic(line)
			}
			var pe PinElem
			pe.Pin = pa[0]
			pe.Acc = pa[1]
			pes = append(pes, pe)
		}
		if len(pes) == 0 {
			panic(line)
		}
		elems[k] = pes
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return
}
