package pongo2trans

var translator ITranslator

type ITranslator interface {
	Translate(in string, args map[string]interface{}) (out string, ok bool)
}

func RegisterTranslator(trans ITranslator) {
	translator = trans
}
