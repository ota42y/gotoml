package gotoml

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/ota42y/gotoml/example"
)

func isValidToml(t *testing.T, b bool) {
	if !b {
		t.Error("invalid toml")
	}
}

func TestExampleData(t *testing.T) {
	// example/normal.go is valid toml struct?
	var n example.Normal

	file, err := os.Open("example/normal.toml")
	if err != nil {
		t.Error(err)
	}
	// read toml file
	if _, err := toml.DecodeReader(file, &n); err != nil {
		t.Error(err)
	}

	isValidToml(t, n.Name == "Dream Sensation")

	d, err := time.Parse("2006-01-02T15:04:05Z", "2015-01-31T18:08:03Z")
	if err != nil {
		t.Error(err)
	}
	isValidToml(t, n.Date == d)

	dates := make([]time.Time, 2)
	dates[0], err = time.Parse("2006-01-02T15:04:05Z", "2015-01-31T18:08:03Z")
	if err != nil {
		t.Error(err)
	}
	dates[1], err = time.Parse("2006-01-02T15:04:05Z", "2015-02-01T16:08:03Z")
	if err != nil {
		t.Error(err)
	}

	isValidToml(t, n.Dates[0] == dates[0])
	isValidToml(t, n.Dates[1] == dates[1])

	isValidToml(t, n.Num == 64)

	isValidToml(t, n.Numbers[0] == 1)
	isValidToml(t, n.Numbers[1] == 2)
	isValidToml(t, n.Numbers[2] == 3)

	isValidToml(t, n.Group.Name == "μ's")
	isValidToml(t, n.Group.School == "National Society of Otonoki")

	isValidToml(t, n.Site.Name == "Saitama Super Arena")
	isValidToml(t, n.Site.Num == 36500)
}

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
	pkgName := "example"

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

func checkGetTypeName(input string, expectName string, expectPackage string, t *testing.T) *generator {
	g := newGenerator()

	var data map[string]interface{}
	if _, err := toml.DecodeReader(strings.NewReader(input), &data); err != nil {
		t.Error(err)
		return nil
	}

	key := ""
	for k := range data {
		key = k
	}

	actual, pkgName := g.getTypeName(key, data[key])
	if expectName != actual {
		t.Errorf("getTypeName should return name %s but `%s`", expectName, actual)
	}
	if pkgName != expectPackage {
		t.Errorf("getTypeName should return package name %s but `%s`", expectPackage, pkgName)
	}

	return g
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

	input = `[group]
name = "A-RISE"
school = "Society of UTX"
`
	g := checkGetTypeName(input, "Group", "", t)
	_, ok := g.tomlData["Group"]
	if !ok {
		t.Errorf("getTypeName should save sub struct but not saved %v", g.tomlData)
	}

	v, ok := g.tomlParsed["Group"]
	if !ok || v {
		t.Errorf("getTypeName should save sub struct but not saved %v", g.tomlParsed)
	}
}
