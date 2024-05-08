package test

import (
	"bytes"
	"github.com/bfun/nesbconvertpin"
	"testing"
)

func TestCSMP_PIN_PARA(t *testing.T) {
	expects:=[][]byte{[]byte("_SVR"),[]byte("_CLT")}
	buf := nesbconvertpin.CSMP_PIN_PARA()
	for _,v:=range expects{
		if !bytes.Contains(buf, v){
			t.Errorf("not contains %v",v)
		}
	}
}
