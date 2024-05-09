package nesbconvertpin

import (
	"fmt"
	"testing"
)

func TestParseAllFormatXml(t *testing.T) {
	var SubExpr bool
	m := ParseAllFormatXml()
	for kf, vf := range m {
		for _, vi := range vf.Items {
			if vi.SubExpr != "" {
				SubExpr = true
				fmt.Println(kf, vi.SubExpr)
			}
		}
	}
	if !SubExpr {
		t.Error("SubExpr not used in ParseAllFormatXml()")
	}
}
