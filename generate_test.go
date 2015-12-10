package gotoml

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestGenerateHead(t *testing.T) {
	pkgName := "main"
	var packages []string

	w := new(bytes.Buffer)
	generateHead(w, pkgName, packages)

	expect := `package main

`
	actual := string(w.Bytes())

	if expect != actual {
		t.Errorf("generateHead expect %s but %s", expect, actual)
	}

	packages = append(packages, "")
	packages = append(packages, "time")
	w = new(bytes.Buffer)
	generateHead(w, pkgName, packages)

	expect = `package main
import (
"time"
)

`
	actual = string(w.Bytes())

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
	if err != nil {
		t.Error(err)
	}

	actual := string(generateByte)
	if expect != actual {
		t.Errorf("Generate expect %s but %s", expect, actual)
	}
}

func TestGetTypeName(t *testing.T) {
	var str string
	expect := "string"
	expectPackage := ""
	actual, pkgName := getTypeName(str)
	if expect != actual {
		t.Errorf("getTypeName(string) should return %s but %s", expect, actual)
	}
	if pkgName != expectPackage {
		t.Errorf("getTypeName(string) should return package name %s but %s", expectPackage, pkgName)
	}

	var dateTime time.Time
	expect = "time.Time"
	expectPackage = "time"
	actual, pkgName = getTypeName(dateTime)
	if expect != actual {
		t.Errorf("getTypeName(time.Time) should return %s but %s", expect, actual)
	}
	if pkgName != expectPackage {
		t.Errorf("getTypeName(string) should return package name %s but %s", expectPackage, pkgName)
	}
}
