package nesbconvertpin

import (
	"encoding/xml"
	"path"
	"regexp"
	"strings"
)

type ServiceTab struct {
	XMLName  xml.Name  `xml:"serviceTab"`
	Services []Service `xml:"service"`
}
type Service struct {
	Name          string `xml:"Name,attr"`
	IFmt          string `xml:"IFmt,attr"`
	EvtIfmtEnd    string `xml:"EvtIfmtEnd"`
	EvtAcallBegin string `xml:"EvtAcallBegin"`
	IsConvertPin  bool
	PinElems      []PinElem
	Matched       []PinElem
	TcElems       []string
	By            string
}

func (i *Service) Clone() Service {
	var s Service
	s.Name = i.Name
	s.IFmt = i.IFmt
	s.EvtIfmtEnd = i.EvtIfmtEnd
	s.EvtAcallBegin = i.EvtAcallBegin
	s.IsConvertPin = i.IsConvertPin
	s.PinElems = make([]PinElem, len(i.PinElems))
	for _, v := range i.PinElems {
		s.PinElems = append(s.PinElems, v)
	}
	s.Matched = make([]PinElem, len(i.Matched))
	for _, v := range i.Matched {
		s.Matched = append(s.Matched, v)
	}
	s.TcElems = make([]string, len(i.TcElems))
	for _, v := range i.TcElems {
		s.TcElems = append(s.TcElems, v)
	}
	s.By = i.By
	return s
}

func trimServiceCDATA(st *ServiceTab) {
	tagdata := regexp.MustCompile(`nesb_get_tagdata\(.*?, *"(.*?)"\)`)
	xmlsign := regexp.MustCompile(`cbs_get_data_by_xmlsign\(.*?,.*?, *"(.*?)"\)`)
	for i, v := range st.Services {
		e := strings.TrimSpace(v.EvtIfmtEnd)
		st.Services[i].EvtIfmtEnd = e
		if strings.Contains(e, "nesbConvertPin") {
			st.Services[i].IsConvertPin = true
		}
		// nesb_get_tagdata("__PACKDATA", "stdpriacno|stdpindata")
		if strings.Contains(e, "nesb_get_tagdata") {
			s := tagdata.FindStringSubmatch(e)
			if len(s) != 2 {
				panic(v)
			}
			st.Services[i].TcElems = strings.Split(s[1], "|")
		}
		// cbs_get_data_by_xmlsign("0", "__PACKDATA", "stdadddtap|stdpriacno|stdpindata")
		if strings.Contains(e, "cbs_get_data_by_xmlsign") {
			s := xmlsign.FindStringSubmatch(e)
			if len(s) != 2 {
				panic(v)
			}
			st.Services[i].TcElems = strings.Split(s[1], "|")
		}
	}
}

func serviceArrayToMap(services []Service) map[string]Service {
	m := make(map[string]Service)
	for _, v := range services {
		m[v.Name] = v
	}
	return m
}

func parseOneServiceXml(fileName string) (dtaName string, services map[string]Service) {
	fullPath := path.Join(getRootDir(), fileName)
	s := strings.TrimSuffix(fileName, "/Service.xml")
	i := strings.LastIndex(s, "/")
	dtaName = s[i+1:]

	var v ServiceTab
	decoder := getGbFileDecoder(fullPath)
	err := decoder.Decode(&v)
	if err != nil {
		panic(err)
	}
	trimServiceCDATA(&v)
	services = serviceArrayToMap(v.Services)
	return
}

func parseAllServiceXml() map[string]map[string]Service {
	m := make(map[string]map[string]Service)
	files := getServiceFiles()
	for _, file := range files {
		dtaName, services := parseOneServiceXml(file)
		m[dtaName] = services
	}
	return m
}
