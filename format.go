package nesbconvertpin

import (
	"encoding/xml"
	"path"
	"regexp"
	"strings"
)

type FormatTab struct {
	XMLName xml.Name `xml:"FormatTab"`
	Formats []Format `xml:"Format"`
}
type Format struct {
	FmtName string       `xml:"FmtName,attr"`
	FmtType string       `xml:"FmtType,attr"`
	Items   []FormatItem `xml:"ItemTab>Item"`
	SubFmts []string
}
type FormatItem struct {
	ItemType string `xml:"ItemType,attr"`
	ItemIgnr string `xml:"ItemIgnr,attr"`
	ElemName string `xml:"ElemName,attr"`
	XmlType  string `xml:"XmlType,attr"`
	XmlName  string `xml:"XmlName,attr"`
	SubName  string `xml:"SubName,attr"`
	SubExpr  string `xml:"SubExpr"`
}

func trimFormatCDATA(formats map[string]Format) {
	re2 := regexp.MustCompile(`.*?\?(.*?):(.*)`)
	re3 := regexp.MustCompile(`.*?\?(.*?):(.*):(.*)`)
	for kf, vf := range formats {
		for ki, vi := range vf.Items {
			if vi.ItemType == "fmt" {
				if vi.SubName == "" {
					panic(kf + "/fmt/SubName empty")
				}
				vf.SubFmts = append(vf.SubFmts, vi.SubName)
			}
			s := strings.TrimSpace(vi.SubExpr)
			if s != "" && strings.Contains(s, "?") && strings.Contains(s, ":") {
				fs := re3.FindStringSubmatch(s)
				if len(fs) > 0 {
					if len(fs) != 4 {
						panic(kf + s)
					}
				} else {
					fs = re2.FindStringSubmatch(s)
					if len(fs) != 3 {
						panic(kf + s)
					}
				}
				for _, f := range fs[1:] {
					f = strings.Replace(f, " ", "", -1)
					f = strings.Replace(f, `"`, "", -1)
					vf.SubFmts = append(vf.SubFmts, f)
				}
			}
			vf.Items[ki].SubExpr = s
		}
		formats[kf] = vf
	}
}

func formatArrayToMap(formats []Format, m map[string]Format) {
	for _, v := range formats {
		m[v.FmtName] = v
	}
}
func parseOneFormatXml(fileName string, m map[string]Format) {
	fullPath := path.Join(getRootDir(), fileName)
	decoder := getGbFileDecoder(fullPath)
	var v FormatTab
	err := decoder.Decode(&v)
	if err != nil {
		panic(err)
	}
	formatArrayToMap(v.Formats, m)
}

func ParseAllFormatXml() map[string]Format {
	m := make(map[string]Format)
	files := getFormatFiles()
	for _, f := range files {
		parseOneFormatXml(f, m)
	}
	trimFormatCDATA(m)
	return m
}
