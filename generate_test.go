package gotoml

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
)

func TestGenerateHead(t *testing.T) {
	pkgName := "main"

	w := new(bytes.Buffer)
	g := newGenerator()
	g.generateHead(w, pkgName)

	expect := `package main

`
	actual := string(w.Bytes())

	if expect != actual {
		t.Errorf("generateHead expect %s but %s", expect, actual)
	}

	g = newGenerator()
	g.usingPackages[""] = true
	g.usingPackages["time"] = true
	w = new(bytes.Buffer)
	g.generateHead(w, pkgName)

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

func checkGetTypeName(input string, expectName string, expectPackage string, t *testing.T) {
	var data map[string]interface{}
	if _, err := toml.DecodeReader(strings.NewReader(input), &data); err != nil {
		t.Error(err)
		return
	}

	key := ""
	for k := range data {
		key = k
	}

	actual, pkgName := getTypeName(data[key])
	if expectName != actual {
		t.Errorf("getTypeName should return name %s but `%s`", expectName, actual)
	}
	if pkgName != expectPackage {
		t.Errorf("getTypeName should return package name %s but `%s`", expectPackage, pkgName)
	}
}

func TestGetTypeName(t *testing.T) {
	input := `name = "Name"`
	checkGetTypeName(input, "string", "", t)

	input = `date = 1979-05-27T07:32:00Z`
	checkGetTypeName(input, "time.Time", "time", t)

	input = `num = 12`
	checkGetTypeName(input, "int64", "", t)

	input = `interface = []`
	checkGetTypeName(input, "[]interface{}", "", t)

	input = `numbers = [1, 2, 3]`
	checkGetTypeName(input, "[]int64", "", t)

	input = `dates = [2015-01-31T18:08:03Z, 2015-02-01T16:08:03Z]`
	checkGetTypeName(input, "[]time.Time", "time", t)
}
