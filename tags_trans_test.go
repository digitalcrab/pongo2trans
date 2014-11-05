package pongo2trans

import (
	pongo2 "gopkg.in/flosch/pongo2.v3"
	"testing"
)

func TestTransNoArgsNoTranslator(t *testing.T) {
	resetTrans()

	var (
		tpl *pongo2.Template
		err error
		res string
	)

	expected := "This is the title"
	source := "{% trans \"This is the title\" %}"

	if tpl, err = pongo2.FromString(source); err != nil {
		t.Fatalf("Error while parsing from string %q: %v", source, err)
	} else if res, err = tpl.Execute(pongo2.Context{}); err != nil {
		t.Fatalf("Error while executing template %q: %v", source, err)
	} else if res != expected {
		t.Errorf("Error in templkate result. Expected %q, but got %q", expected, res)
	}
}

func TestTransNoArgs(t *testing.T) {
	resetTrans()
	registerTrans()

	var (
		tpl *pongo2.Template
		err error
		res string
	)

	expected := "Это заголовок"
	source := "{% trans \"This is the title\" %}"

	if tpl, err = pongo2.FromString(source); err != nil {
		t.Fatalf("Error while parsing from string %q: %v", source, err)
	} else if res, err = tpl.Execute(pongo2.Context{}); err != nil {
		t.Fatalf("Error while executing template %q: %v", source, err)
	} else if res != expected {
		t.Errorf("Error in templkate result. Expected %q, but got %q", expected, res)
	}
}

func TestTransNoTranslation(t *testing.T) {
	resetTrans()
	registerTrans()

	var (
		tpl *pongo2.Template
		err error
		res string
	)

	expected := "end point"
	source := "{% trans \"end point\" %}"

	if tpl, err = pongo2.FromString(source); err != nil {
		t.Fatalf("Error while parsing from string %q: %v", source, err)
	} else if res, err = tpl.Execute(pongo2.Context{}); err != nil {
		t.Fatalf("Error while executing template %q: %v", source, err)
	} else if res != expected {
		t.Errorf("Error in templkate result. Expected %q, but got %q", expected, res)
	}
}

func TestTransVar(t *testing.T) {
	resetTrans()
	registerTrans()

	var (
		tpl *pongo2.Template
		err error
		res string
	)

	ctx := pongo2.Context{
		"v": "This is the title",
	}

	expected := "Это заголовок"
	source := "{% trans v %}"

	if tpl, err = pongo2.FromString(source); err != nil {
		t.Fatalf("Error while parsing from string %q: %v", source, err)
	} else if res, err = tpl.Execute(ctx); err != nil {
		t.Fatalf("Error while executing template %q: %v", source, err)
	} else if res != expected {
		t.Errorf("Error in templkate result. Expected %q, but got %q", expected, res)
	}
}

func TestTransAsVar(t *testing.T) {
	resetTrans()
	registerTrans()

	var (
		tpl *pongo2.Template
		err error
		res string
	)

	expected := "<title>Это заголовок</title>"
	source := "{% trans \"This is the title\" as title %}<title>{{ title }}</title>"

	if tpl, err = pongo2.FromString(source); err != nil {
		t.Fatalf("Error while parsing from string %q: %v", source, err)
	} else if res, err = tpl.Execute(pongo2.Context{}); err != nil {
		t.Fatalf("Error while executing template %q: %v", source, err)
	} else if res != expected {
		t.Errorf("Error in templkate result. Expected %q, but got %q", expected, res)
	}
}

func TestTransAsVarBlock(t *testing.T) {
	resetTrans()
	registerTrans()

	var (
		tpl *pongo2.Template
		err error
		res string
	)

	expected := "Это заголовок"
	source := "{% block main %}{% trans \"This is the title\" as title %}{% endblock %}{{ title }}"

	if tpl, err = pongo2.FromString(source); err != nil {
		t.Fatalf("Error while parsing from string %q: %v", source, err)
	} else if res, err = tpl.Execute(pongo2.Context{}); err != nil {
		t.Fatalf("Error while executing template %q: %v", source, err)
	} else if res != expected {
		t.Errorf("Error in templkate result. Expected %q, but got %q", expected, res)
	}
}
