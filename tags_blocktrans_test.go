package pongo2trans

import (
	"github.com/flosch/pongo2"
	"testing"
)

func TestTransBlockNoArgsNoTranslator(t *testing.T) {
	resetTrans()

	var (
		tpl	*pongo2.Template
		err	error
		res	string
	)

	expected := "This string will have value inside."
	source := "{% blocktrans %}This string will have value inside.{% endblocktrans %}"

	if tpl, err = pongo2.FromString(source); err != nil {
		t.Fatalf("Error while parsing from string %q: %v", source, err)
	} else if res, err = tpl.Execute(pongo2.Context{}); err != nil {
		t.Fatalf("Error while executing template %q: %v", source, err)
	} else if res != expected {
		t.Errorf("Error in templkate result. Expected %q, but got %q", expected, res)
	}
}

func TestTransBlockNoArgs(t *testing.T) {
	resetTrans()
	registerTrans()

	var (
		tpl	*pongo2.Template
		err	error
		res	string
	)

	expected := "Это заголовок"
	source := "{% blocktrans %}This is the title{% endblocktrans %}"

	if tpl, err = pongo2.FromString(source); err != nil {
		t.Fatalf("Error while parsing from string %q: %v", source, err)
	} else if res, err = tpl.Execute(pongo2.Context{}); err != nil {
		t.Fatalf("Error while executing template %q: %v", source, err)
	} else if res != expected {
		t.Errorf("Error in templkate result. Expected %q, but got %q", expected, res)
	}
}

func TestTransBlockVar(t *testing.T) {
	resetTrans()
	registerTrans()

	var (
		tpl	*pongo2.Template
		err	error
		res	string
	)

	ctx := pongo2.Context{
		"v": "This is the title",
	}

	expected := "Это заголовок"
	source := "{% blocktrans %}{{ v }}{% endblocktrans %}"

	if tpl, err = pongo2.FromString(source); err != nil {
		t.Fatalf("Error while parsing from string %q: %v", source, err)
	} else if res, err = tpl.Execute(ctx); err != nil {
		t.Fatalf("Error while executing template %q: %v", source, err)
	} else if res != expected {
		t.Errorf("Error in templkate result. Expected %q, but got %q", expected, res)
	}
}

func TestTransBlockWith(t *testing.T) {
	resetTrans()
	registerTrans()

	var (
		tpl	*pongo2.Template
		err	error
		res	string
	)

	expected := "Переведи меня: Я перевод"
	source := "{% blocktrans with \"%v%\"=\"Я перевод\" %}Translate me: %v%{% endblocktrans %}"

	if tpl, err = pongo2.FromString(source); err != nil {
		t.Fatalf("Error while parsing from string %q: %v", source, err)
	} else if res, err = tpl.Execute(pongo2.Context{}); err != nil {
		t.Fatalf("Error while executing template %q: %v", source, err)
	} else if res != expected {
		t.Errorf("Error in templkate result. Expected %q, but got %q", expected, res)
	}
}
