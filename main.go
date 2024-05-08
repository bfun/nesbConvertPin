package nesbconvertpin

func main() {
	mcps := CSMP_PIN_SERVICE()
	gets := get_svcname_by_procode()
	txml := nesb_txml()
	mpes := CSMP_PIN_ELEM()
	mfmt := parseAllFormatXml()
	mdta := getAllConvertPinDtas()
	meshFmt(mdta, mfmt, mpes)
	meshTxml(mdta, txml, mpes)
	meshGets(mdta, gets, mpes)
	patchJSON_SVR(mdta, mpes)
	patchCSMP_PIN_SERVICE(mdta, mcps)
	writeResult(mdta)
}
