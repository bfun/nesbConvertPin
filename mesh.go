package nesbconvertpin

import (
	"fmt"
	"strings"
)

func getPinElemsByService(dta string, svc string, mpes map[string][]PinElem) []PinElem {
	for k, v := range mpes {
		ds := strings.Split(k, ".")
		n := len(ds)
		if n < 2 {
			continue
		}
		if ds[0] != dta {
			continue
		}
		last := ds[n-1]
		if last == "01" || last == "02" {
			continue
		}
		if strings.Contains(svc, last) || strings.Contains(last, svc) {
			return v
		}
	}
	var pes []PinElem
	for k, v := range mpes {
		if k == dta || k == dta+".01" || k == dta+".02" {
			pes = append(pes, v...)
		}
	}
	return pes
}
func meshFmt(dtas map[string]DataTransferAdapter, fmts map[string]Format, mpes map[string][]PinElem) {
	for kd, vd := range dtas {
		for ks, vs := range vd.Services {
			if !vd.ConvertPin && !vs.ConvertPin {
				continue
			}
			pes := getPinElemsByService(kd, ks, mpes)
			if len(pes) == 0 {
				fmt.Printf("??? %v.%v DTA_%v/Service_%v, but no rules matched\n", kd, ks, vd.ConvertPin, vs.ConvertPin)
				continue
			}
			var elems []PinElem
			f, ok := fmts[vs.IFmt]
			if !ok {
				panic(vs.IFmt)
			}
			for _, pe := range pes {
				var pin, acc bool
				for _, vi := range f.Items {
					if vi.ElemName == pe.Pin {
						pin = true
					}
				}
				for _, vi := range f.Items {
					if vi.ElemName == pe.Acc {
						acc = true
					}
				}
				if pin && acc {
					elems = append(elems, pe)
				}
			}
			if len(elems) == 0 {
				continue
			}
			vs.Matched = elems
			vd.Services[ks] = vs
			dtas[kd] = vd
			fmt.Printf("meshFmt matched: %v.%v DTA_%v/Service_%v fmt[%v] elems[%v]\n", kd, ks, vd.ConvertPin, vs.ConvertPin, vs.IFmt, elems)
		}
	}
}

func meshTxml(dtas map[string]DataTransferAdapter, txml map[string]map[string]NesbTxml, mpes map[string][]PinElem) {
	for kd, vd := range txml {
		dta, ok := dtas[kd]
		if !ok {
			fmt.Printf("nesb_txml[%v] not in dtas-by-xml\n", kd)
			continue
		}
		for ks, vs := range vd {
			svc, ok := dta.Services[ks]
			if !ok {
				svc = Service{}
				dta.Services[ks] = svc
			}
			if !dta.ConvertPin && !svc.ConvertPin {
				continue
			}
			pes := getPinElemsByService(kd, ks, mpes)
			if len(pes) == 0 {
				fmt.Printf("??? nesb_txml %v.%v DTA_%v/Service_%v, but no rules matched\n", kd, ks, dta.ConvertPin, svc.ConvertPin)
				continue
			}
			var matched []PinElem
			var elems []string
			if len(vs.Elems) > 0 {
				elems = vs.Elems
			}
			if len(svc.TcElems) > 0 {
				elems = append(elems, svc.TcElems...)
			}
			if len(elems) == 0 {
				continue
			}
			for _, pe := range pes {
				var pin, acc bool
				for _, v := range elems {
					if v == pe.Pin {
						pin = true
					}
				}
				for _, v := range elems {
					if v == pe.Acc {
						acc = true
					}
				}
				if pin && acc {
					matched = append(matched, pe)
				}
			}
			if len(matched) == 0 {
				continue
			}
			svc2 := svc.Clone()
			svc2.Matched = matched
			svc2.By = "nesb_txml.txt"
			dta.Services[ks] = svc2
			fmt.Printf("meshTxml matched: %v.%v DTA_%v/Service_%v %v elems[%v]\n", kd, ks, dta.ConvertPin, svc.ConvertPin, svc2.By, elems)
		}
		dtas[kd] = dta
	}
}

func meshGets(dtas map[string]DataTransferAdapter, gets map[string]map[string][]string, mpes map[string][]PinElem) {
	for kd, vd := range gets {
		dta, ok := dtas[kd]
		if !ok {
			fmt.Printf("get_svcname_by_procode[%v] not in dtas-by-xml\n", kd)
			continue
		}
		for ks, vs := range vd {
			svc, ok := dta.Services[ks]
			if !ok {
				svc = Service{}
				dta.Services[ks] = svc
			}
			if !dta.ConvertPin && !svc.ConvertPin {
				continue
			}
			pes := getPinElemsByService(kd, ks, mpes)
			if len(pes) == 0 {
				fmt.Printf("??? get_svcname_by_procode %v.%v DTA_%v/Service_%v, but no rules matched\n", kd, ks, dta.ConvertPin, svc.ConvertPin)
				continue
			}
			var matched []PinElem
			var elems []string
			if len(svc.TcElems) > 0 {
				elems = svc.TcElems
			}
			if len(elems) == 0 {
				continue
			}
			for _, pe := range pes {
				var pin, acc bool
				for _, v := range elems {
					if v == pe.Pin {
						pin = true
					}
				}
				for _, v := range elems {
					if v == pe.Acc {
						acc = true
					}
				}
				if pin && acc {
					matched = append(matched, pe)
				}
			}
			if len(matched) == 0 {
				continue
			}
			for _, cod := range vs {
				svc2 := svc.Clone()
				svc2.Matched = matched
				svc2.By = "get_svcname_by_procode.txt"
				dta.Services[cod] = svc2
				fmt.Printf("meshGets matched: %v.%v DTA_%v/Service_%v %v elems[%v]\n", kd, ks, dta.ConvertPin, svc.ConvertPin, svc2.By, elems)
			}
		}
		dtas[kd] = dta
	}
}

func patchJSON_SVR(dtas map[string]DataTransferAdapter, mpes map[string][]PinElem) {
	dtaNames := []string{"JSON_SVR", "JSON1_SVR"}
	for _, dtaName := range dtaNames {
		var dta DataTransferAdapter
		dta.Services = make(map[string]Service)
		for k, v := range mpes {
			if !strings.HasPrefix(k, dtaName) {
				continue
			}
			ds := strings.Split(k, ".")
			if len(ds) < 2 {
				continue
			}
			var s Service
			s.Name = ds[1]
			s.ConvertPin = true
			s.PinElems = v[:]
			s.Matched = v[:]
			s.By = "CSMP_PIN_ELEM.txt"
			dta.Services[s.Name] = s
		}
		dtas[dtaName] = dta
	}
}

func patchCSMP_PIN_SERVICE(dtas map[string]DataTransferAdapter, mcps map[string]map[string][]PinElem) {
	i := 0
	for kd, vd := range mcps {
		dta, ok := dtas[kd]
		if !ok {
			panic(kd)
		}
		for ks, vs := range vd {
			var s Service
			s.Name = ks
			s.ConvertPin = true
			s.PinElems = vs[:]
			s.Matched = vs[:]
			s.By = "CSMP_PIN_SERVICE.txt"
			dta.Services[s.Name] = s
			i++
			fmt.Printf("patchCSMP_PIN_SERVICE: %v %v %v\n", i, kd, ks)
		}
		dtas[kd] = dta
	}
}

func trimPOBS(dtas map[string]DataTransferAdapter) {
	target := "POBS_SVR"
	dta, ok := dtas[target]
	if !ok {
		dta = DataTransferAdapter{}
	}
	svcs := make(map[string]Service)
	for kd, vd := range dtas {
		if !strings.Contains(kd, "POBS") || !strings.Contains(kd, "_SVR") {
			continue
		}
		for ks, vs := range vd.Services {
			if vd.ConvertPin {
				vs.ConvertPin = true
			}
			if vs.ConvertPin && len(vs.Matched) > 0 {
				svcs[ks] = vs
			}
		}
	}
	dta.Services = svcs
	dtas[target] = dta
}
