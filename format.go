package nesbconvertpin

import (
	"encoding/xml"
	"path"
)

type FormatTab struct {
	XMLName xml.Name `xml:"FormatTab"`
	Formats []Format `xml:"Format"`
}
type Format struct {
	FmtName string       `xml:"FmtName,attr"`
	FmtType string       `xml:"FmtType,attr"`
	Items   []FormatItem `xml:"ItemTab>Item"`
}
type FormatItem struct {
	ItemType string `xml:"ItemType,attr"`
	ItemIgnr string `xml:"ItemIgnr,attr"`
	ElemName string `xml:"ElemName,attr"`
	XmlType  string `xml:"XmlType,attr"`
	XmlName  string `xml:"XmlName,attr"`
	SubName  string `xml:"SubName,attr"`
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

func parseAllFormatXml() map[string]Format {
	m := make(map[string]Format)
	files := getFormatFiles()
	for _, f := range files {
		parseOneFormatXml(f, m)
	}
	return m
}
