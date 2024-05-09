package test

import (
	"fmt"
	"github.com/bfun/nesbconvertpin"
	"testing"
)

func TestParseAllDtaParmXml(t *testing.T) {
	m := nesbconvertpin.ParseAllDtaParmXml()
	var Name, EvtIprtcfmtBegin, EvtIprtcfmtEnd, EvtIfmtEnd, EvtOfmtBegin, EvtOprtcfmtBegin, ConvertPin, Services, NESB_SDTA_NAME bool
	for k, v := range m {
		fmt.Println(k, v.NESB_SDTA_NAME)
		if v.Name != "" {
			Name = true
		}
		if v.EvtIprtcfmtBegin != "" {
			EvtIprtcfmtBegin = true
		}
		if v.EvtIprtcfmtEnd != "" {
			EvtIprtcfmtEnd = true
		}
		if v.EvtIfmtEnd != "" {
			EvtIfmtEnd = true
		}
		if v.EvtOfmtBegin != "" {
			EvtOfmtBegin = true
		}
		if v.EvtOprtcfmtBegin != "" {
			EvtOprtcfmtBegin = true
		}
		if v.ConvertPin {
			ConvertPin = true
		}
		if len(v.Services) > 0 {
			Services = true
		}
		if v.NESB_SDTA_NAME != "" {
			NESB_SDTA_NAME = true
		}
	}
	if !Name {
		t.Error("Name is missing")
	}
	if !EvtIprtcfmtBegin {
		t.Error("EvtIprtcfmtBegin is missing")
	}
	if !EvtIprtcfmtEnd {
		t.Error("EvtIprtcfmtEnd is missing")
	}
	if !EvtIfmtEnd {
		t.Error("EvtIfmtEnd is missing")
	}
	if !EvtOfmtBegin {
		t.Error("EvtOfmtBegin is missing")
	}
	if !EvtOprtcfmtBegin {
		t.Error("EvtOprtcfmtBegin is missing")
	}
	if !ConvertPin {
		t.Error("ConvertPin is missing")
	}
	if !Services {
		// t.Error("Services is missing")
	}
	if !NESB_SDTA_NAME {
		t.Error("NESB_SDTA_NAME is missing")
	}
}
