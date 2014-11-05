package pongo2trans

import (
	"fmt"
	"strings"
	"testing"
)

type testTranslator struct {
	data map[string]string
}

func (self *testTranslator) Translate(in string, args map[string]interface{}) (string, bool) {
	res, ok := self.data[in]
	if !ok {
		return in, false
	}

	if args != nil && len(args) > 0 {
		for key, value := range args {
			res = strings.Replace(res, key, fmt.Sprintf("%v", value), -1)
		}
	}

	return res, true
}

func registerTrans() *testTranslator {
	trans := &testTranslator{
		data: map[string]string{
			"This is the title": "Это заголовок",
			"Translate me: %v%": "Переведи меня: %v%",
		},
	}
	RegisterTranslator(trans)
	return trans
}

func resetTrans() {
	RegisterTranslator(nil)
}

func TestRegisterTranslator(t *testing.T) {
	resetTrans()
	trans := registerTrans()
	if translator != trans {
		t.Errorf("Translator is not registered")
	}
}
