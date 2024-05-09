package test

import (
	"fmt"
	"github.com/bfun/nesbconvertpin"
	"testing"
)

func Test_parseAllServiceXml(t *testing.T) {
	var Name, IFmt, EvtIfmtBegin, EvtIfmtEnd, EvtAcallBegin, NESB_SDTA_NAME, ConvertPin, TcElems bool
	m := nesbconvertpin.ParseAllServiceXml()
	for kd, vd := range m {
		for ks, vs := range vd {
			var s []string
			if vs.Name != "" {
				Name = true
			}
			if vs.IFmt != "" {
				IFmt = true
			}
			if vs.EvtIfmtBegin != "" {
				EvtIfmtBegin = true
				s = append(s, vs.EvtIfmtBegin)
			}
			if vs.EvtIfmtEnd != "" {
				EvtIfmtEnd = true
				s = append(s, vs.EvtIfmtEnd)
			}
			if vs.EvtAcallBegin != "" {
				EvtAcallBegin = true
				s = append(s, vs.EvtAcallBegin)
			}
			if vs.NESB_SDTA_NAME != "" {
				NESB_SDTA_NAME = true
				s = append(s, vs.NESB_SDTA_NAME)
			}
			if vs.ConvertPin {
				ConvertPin = true
			}
			if len(vs.TcElems) > 0 {
				TcElems = true
				s = append(s, vs.TcElems...)
			}
			if len(s) > 0 {
				fmt.Println(kd, ks, s)
			}
		}
	}
	if !Name {
		t.Error("Name not used in ParseAllServiceXml")
	}
	if !IFmt {
		t.Error("IFmt not used in ParseAllServiceXml")
	}
	if !EvtIfmtBegin {
		t.Error("EvtIfmtBegin not used in ParseAllServiceXml")
	}
	if !EvtIfmtEnd {
		t.Error("EvtIfmtEnd not used in ParseAllServiceXml")
	}
	if !EvtAcallBegin {
		t.Error("EvtAcallBegin not used in ParseAllServiceXml")
	}
	if !NESB_SDTA_NAME {
		t.Error("NESB_SDTA_NAME not used in ParseAllServiceXml")
	}
	if !ConvertPin {
		t.Error("ConvertPin not used in ParseAllServiceXml")
	}
	if !TcElems {
		t.Error("TcElems not used in ParseAllServiceXml")
	}
}
