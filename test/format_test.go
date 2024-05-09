package test

import (
	"fmt"
	"github.com/bfun/nesbconvertpin"
	"testing"
)

func TestParseAllFormatXml(t *testing.T) {
	var SubFmts, SubExpr bool
	m := nesbconvertpin.ParseAllFormatXml()
	for kf, vf := range m {
		for _, vi := range vf.Items {
			if vi.SubExpr != "" {
				SubExpr = true
				fmt.Println(kf, vi.SubExpr)
			}
		}
		if len(vf.SubFmts) > 0 {
			SubFmts = true
			fmt.Println(kf, vf.SubFmts)
		}
	}
	if !SubFmts {
		t.Error("SubFmts not used in ParseAllFormatXml()")
	}
	if !SubExpr {
		t.Error("SubExpr not used in ParseAllFormatXml()")
	}
}
