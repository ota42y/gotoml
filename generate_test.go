package gotoml

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerateHead(t *testing.T) {
	structName := "Foo"
	pkgName := "main"

	w := new(bytes.Buffer)
	generateHead(w, structName, pkgName)

	expect := `package main

type Foo struct {
`
	actual := string(w.Bytes())

	if expect != actual {
		t.Errorf("generateHead expect %s but %s", expect, actual)
	}
}

func TestGenerate(t *testing.T) {
	structName := "Normal"
	pkgName := "normal"

	file, err := os.Open("example/normal.toml")
	if err != nil {
		t.Error(err)
	}

	b, err := ioutil.ReadFile("example/normal.go")
	if err != nil {
		t.Error(err)
	}
	expect := string(b)

	generateByte, err := Generate(file, structName, pkgName)
	actual := string(generateByte)
	if expect != actual {
		t.Errorf("Generate expect %s but %s", expect, actual)
	}
}
