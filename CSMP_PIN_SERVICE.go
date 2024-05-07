package nesbconvertpin

import (
	"bufio"
	"os"
	"path"
	"strings"
)

func CSMP_PIN_SERVICE() (services map[string]map[string][]PinElem) {
	services = make(map[string]map[string][]PinElem)
	fileName := path.Join(getRootDir(), "etc/enum/CSMP_PIN_SERVICE.txt")
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "enumData1") {
			continue
		}
		kv := strings.SplitN(line, "\t", 2)
		if len(kv) < 2 {
			continue
		}
		ks := strings.Split(kv[0], ".")
		if len(ks) < 3 {
			panic(line)
		}
		dtaname := ks[0]
		if !strings.Contains(dtaname, "_SVR") && !strings.Contains(dtaname, "_SGET") && !strings.Contains(dtaname, "_PAY") && !strings.Contains(dtaname, "_CLS") {
			continue
		}
		procode := ks[2]
		vs := strings.Split(kv[1], "$")
		if len(vs) < 2 {
			panic(line)
		}
		var pes []PinElem
		couples := strings.Split(vs[1], "|")
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
		dta, ok := services[dtaname]
		if !ok {
			dta = make(map[string][]PinElem)
		}
		dta[procode] = pes
		services[dtaname] = dta
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return
}
