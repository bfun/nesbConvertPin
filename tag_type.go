package nesbconvertpin

import (
	"bufio"
	"os"
	"path"
	"regexp"
	"strings"
)

func tag_type() (tags map[string][]string) {
	tags = make(map[string][]string)
	re := regexp.MustCompile("get_req:{(.*?)}")
	fileName := path.Join(getRootDir(), "etc/enum/tag_type.txt")
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") || !strings.Contains(line, "get_req:{") {
			continue
		}
		kv := strings.Split(line, "\t")
		if len(kv) != 2 {
			continue
		}
		vs := re.FindStringSubmatch(kv[1])
		if len(vs) != 2 {
			panic(line)
		}
		couples := strings.Split(vs[1], ",")
		var elems []string
		for _, c := range couples {
			te := strings.Split(c, ":")
			if len(te) != 2 {
				panic(line)
			}
			elems = append(elems, te[1])
		}
		tags[kv[0]] = elems
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return
}
