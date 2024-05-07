package nesbconvertpin

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"sort"
	"strconv"
	"strings"
)

func writeResult(result map[string]DtaConvertPin) {
	f := excelize.NewFile()
	defer f.Close()
	sheetName := "转加密交易"
	f.SetSheetName("Sheet1", sheetName)
	writeSheet(f, sheetName, result)
	name := "转加密交易清单_v1.2.xlsx"
	if err := f.SaveAs(name); err != nil {
		panic(err)
	}
}

func writeSheet(f *excelize.File, sheetName string, result map[string]DtaConvertPin) {
	style, err := f.NewStyle(&excelize.Style{NumFmt: 1})
	if err != nil {
		panic(err)
	}
	f.SetColWidth(sheetName, "A", "A", 15)
	f.SetColWidth(sheetName, "B", "B", 50)
	f.SetColWidth(sheetName, "C", "C", 100)
	f.SetColWidth(sheetName, "D", "D", 100)
	f.SetCellStr(sheetName, "A1", "DTA")
	f.SetCellStr(sheetName, "B1", "交易码")
	f.SetCellStr(sheetName, "C1", "报文格式")
	f.SetCellStr(sheetName, "D1", "转加密字段")
	i := 0
	var dtas []string
	for k, _ := range result {
		dtas = append(dtas, k)
	}
	sort.Strings(dtas)
	for _, kd := range dtas {
		if strings.Contains(kd, "POBS") && kd != "POBS_SVR" {
			continue
		}
		dta := result[kd]
		var svcs []string
		for k, _ := range dta.Services {
			svcs = append(svcs, k)
		}
		sort.Strings(svcs)
		for _, ks := range svcs {
			svc := dta.Services[ks]
			if len(svc.Matched) == 0 {
				continue
			}
			n := strconv.Itoa(i + 2)
			var s string
			for _, pe := range svc.Matched {
				if s == "" {
					s = fmt.Sprintf("%v,%v", pe.Pin, pe.Acc)
				} else {
					s += fmt.Sprintf("|%v,%v", pe.Pin, pe.Acc)
				}
			}
			f.SetCellStyle(sheetName, "A"+n, "D"+n, style)
			f.SetCellStr(sheetName, "A"+n, kd)
			f.SetCellStr(sheetName, "B"+n, ks)
			if svc.By != "" {
				f.SetCellStr(sheetName, "C"+n, svc.By)
			} else {
				f.SetCellStr(sheetName, "C"+n, svc.IFmt)
			}
			f.SetCellStr(sheetName, "D"+n, s)
			i++
		}
	}
}
