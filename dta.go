package nesbconvertpin

import (
	"encoding/xml"
	"path"
	"strings"
)

type DataTransferAdapter struct {
	XMLName          xml.Name `xml:"DataTransferAdapter"`
	Name             string   `xml:"Name,attr"`
	EvtIprtcfmtBegin string   `xml:"EvtIprtcfmtBegin"`
	EvtIprtcfmtEnd   string   `xml:"EvtIprtcfmtEnd"`
	EvtIfmtEnd       string   `xml:"EvtIfmtEnd"`
	EvtOfmtBegin     string   `xml:"EvtOfmtBegin"`
	EvtOprtcfmtBegin string   `xml:"EvtOprtcfmtBegin"`
}

func parseOneDtaParmXml(fileName string) DataTransferAdapter {
	fullPath := path.Join(getRootDir(), fileName)
	var v DataTransferAdapter
	decoder := getGbFileDecoder(fullPath)
	err := decoder.Decode(&v)
	if err != nil {
		panic(err)
	}
	// trimServiceCDATA(&v)
	return v
}

func ParseAllDtaParmXml() map[string]DataTransferAdapter {
	m := make(map[string]DataTransferAdapter)
	files := getDtaParmFiles()
	for _, file := range files {
		dta := parseOneDtaParmXml(file)
		m[dta.Name] = dta
	}
	return m
}

type DtaConvertPin struct {
	All      bool
	Services map[string]Service
}

func isDtaConvertPin(filepath string) (dtaName string, is bool) {
	s := strings.TrimSuffix(filepath, "/DtaParm.xml")
	i := strings.LastIndex(s, "/")
	dtaName = s[i+1:]
	is = false
	target := "nesbConvertPin"
	file, scanner := fileScanner(filepath)
	defer file.Close()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.Contains(line, target) {
			is = true
			return
		}
		if strings.HasPrefix(line, "<ErrTab") {
			return
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return
}

func getAllConvertPinDtasByDtaParm() map[string]DtaConvertPin {
	m := make(map[string]DtaConvertPin)
	files := getDtaParmFiles()
	for _, file := range files {
		dtaName, is := isDtaConvertPin(file)
		if !is {
			continue
		}
		var dcp DtaConvertPin
		dcp.All = true
		m[dtaName] = dcp
	}
	return m
}

func getAllConvertPinDtas() map[string]DtaConvertPin {
	dtas := getAllConvertPinDtasByDtaParm()
	svcs := parseAllServiceXml()
	for k, v := range svcs {
		d, ok := dtas[k]
		if ok {
			d.Services = v
			dtas[k] = d
		} else {
			var dcp DtaConvertPin
			dcp.All = false
			dcp.Services = v
			dtas[k] = dcp
		}
	}
	return dtas
}
