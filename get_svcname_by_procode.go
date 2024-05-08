package nesbconvertpin

import (
	"bufio"
	"os"
	"path"
	"strings"
)

func get_svcname_by_procode() (services map[string]map[string][]string) {
	services = make(map[string]map[string][]string)
	fileName := path.Join(getRootDir(), "etc/enum/get_svcname_by_procode.txt")
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "enumData1") || strings.HasPrefix(line, "*") {
			continue
		}
		kv := strings.Split(line, "\t")
		if len(kv) != 2 {
			panic(line)
		}
		svc := kv[1]
		if strings.Contains(kv[0], ".") {
			ds := strings.Split(kv[0], ".")
			dta, ok := services[ds[0]]
			if ok {
				codes, ok := dta[svc]
				if ok {
					codes = append(codes, ds[1])
					dta[svc] = codes
				} else {
					dta[svc] = []string{ds[1]}
				}
			} else {
				dta = make(map[string][]string)
				dta[svc] = []string{ds[1]}
			}
			services[ds[0]] = dta
		} else {
			dtaname := "TXML_SVR"
			dta, ok := services[dtaname]
			if ok {
				codes, ok := dta[svc]
				if ok {
					codes = append(codes, kv[0])
					dta[svc] = codes
				} else {
					dta[svc] = []string{kv[0]}
				}
			} else {
				dta = make(map[string][]string)
				dta[svc] = []string{kv[0]}
			}
			services[dtaname] = dta
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return
}
