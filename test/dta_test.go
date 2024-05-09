package test

import (
	"fmt"
	"github.com/bfun/nesbconvertpin"
	"testing"
)

func TestParseAllDtaParmXml(t *testing.T) {
	m := nesbconvertpin.ParseAllDtaParmXml()
	var name, evtIprtcfmtBegin, evtIprtcfmtEnd, evtIfmtEnd, evtOfmtBegin, evtOprtcfmtBegin bool
	for k, v := range m {
		fmt.Println(k, v)
		if v.Name != "" {
			name = true
		}
		if v.EvtIprtcfmtBegin != "" {
			evtIprtcfmtBegin = true
		}
		if v.EvtIprtcfmtEnd != "" {
			evtIprtcfmtEnd = true
		}
		if v.EvtIfmtEnd != "" {
			evtIfmtEnd = true
		}
		if v.EvtOfmtBegin != "" {
			evtOfmtBegin = true
		}
		if v.EvtOprtcfmtBegin != "" {
			evtOprtcfmtBegin = true
		}
	}
	if !name {
		t.Error("Name is missing")
	}
	if !evtIprtcfmtBegin {
		t.Error("EvtIprtcfmtBegin is missing")
	}
	if !evtIprtcfmtEnd {
		t.Error("EvtIprtcfmtEnd is missing")
	}
	if !evtIfmtEnd {
		t.Error("EvtIfmtEnd is missing")
	}
	if !evtOfmtBegin {
		t.Error("EvtOfmtBegin is missing")
	}
	if !evtOprtcfmtBegin {
		t.Error("EvtOprtcfmtBegin is missing")
	}
}
