package nesbconvertpin

import (
	"bufio"
	"os"
	"path"
	"strings"
)

type NesbTxml struct {
	Tag     string
	Service string
	Elems   []string
}

func nesb_txml() (services map[string]map[string]NesbTxml) {
	services = make(map[string]map[string]NesbTxml)
	tags := tag_type()
	fileName := path.Join(getRootDir(), "etc/enum/nesb_txml.txt")
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") || !strings.Contains(line, ".") {
			continue
		}
		kv := strings.Split(line, "\t")
		if len(kv) != 2 {
			continue
		}
		ds := strings.Split(kv[0], ".")
		if len(ds) != 2 {
			continue
		}
		dtaname := ds[0]
		if !strings.Contains(dtaname, "_SVR") && !strings.Contains(dtaname, "_SGET") && !strings.Contains(dtaname, "_PAY") && !strings.Contains(dtaname, "_CLS") {
			continue
		}
		procode := ds[1]
		var nt NesbTxml
		ts := strings.Split(kv[1], "|")
		if len(ts) == 1 {
			dds := strings.Split(ts[0], ".")
			if len(dds) < 2 {
				continue
			}
			nt.Service = dds[1]
		}
		if len(ts) == 2 {
			nt.Tag = ts[0]
			dds := strings.Split(ts[1], ".")
			if len(dds) < 2 {
				continue
			}
			nt.Service = dds[1]
			nt.Elems = tags[nt.Tag]
		}
		if nt.Service == "" {
			continue
		}
		dta, ok := services[dtaname]
		if !ok {
			dta = make(map[string]NesbTxml)
			services[dtaname] = dta
		}
		dta[procode] = nt
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return
}
