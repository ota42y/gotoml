package gotoml

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"reflect"
	"sort"

	"github.com/BurntSushi/toml"
	"strings"
)

func Generate(input io.Reader, structName string, pkgName string) ([]byte, error) {
	g := newGenerator()

	// read toml file
	var tomlMap map[string]interface{}
	if _, err := toml.DecodeReader(input, &tomlMap); err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)
	g.generateBody(body, structName, tomlMap)

	w := new(bytes.Buffer)
	g.generateHead(w, pkgName)
	fmt.Fprintf(w, body.String())

	result, err := format.Source(w.Bytes())
	if err != nil {
		return nil, fmt.Errorf("go format error %s when %s formated", err.Error(), w.String())
	}

	return result, nil
}

type generator struct {
	tomlData      map[string]interface{} // already checked data
	usingPackages map[string]bool
}

func newGenerator() *generator {
	return &generator{
		tomlData:      make(map[string]interface{}),
		usingPackages: make(map[string]bool),
	}
}

func (g *generator) generateHead(w io.Writer, pkgName string) {
	fmt.Fprintf(w, "package %s\n", pkgName)

	// write all package names
	var packages []string
	for k := range g.usingPackages {
		packages = append(packages, k)
	}
	sort.Strings(packages)

	if len(packages) != 0 {
		fmt.Fprintf(w, "import (\n")

		for _, name := range packages {
			if name != "" {
				fmt.Fprintf(w, "\"%s\"\n", name)
			}
		}

		fmt.Fprintf(w, ")\n")
	}

	fmt.Fprintf(w, "\n")
}

func (g *generator) generateBody(w io.Writer, structName string, data map[string]interface{}) {
	fmt.Fprintf(w, "type %s struct {\n", structName)

	// sort keys
	mk := make([]string, len(data))
	i := 0
	for k, _ := range data {
		mk[i] = k
		i++
	}
	sort.Strings(mk)

	// create struct
	for _, key := range mk {
		keyTitle := strings.Title(key)
		typeName, pkgPath := getTypeName(data[key])

		// save package name
		g.usingPackages[pkgPath] = true

		fmt.Fprintf(w, "%s %s `toml:\"%s\"`\n", keyTitle, typeName, key)
	}
	fmt.Fprintf(w, "}")
}

func getTypeName(i interface{}) (string, string) {
	t := reflect.TypeOf(i)
	pkgPath := t.PkgPath()
	// if specific package's struct(like time.Time), return with package name
	if pkgPath != "" {
		return fmt.Sprintf("%s.%s", t.PkgPath(), t.Name()), pkgPath
	}

	switch iType := i.(type) {
	case int:
		return "int64", ""
	case []interface{}:
		if len(iType) == 0 {
			// no items, so we cant't decide type.
			return "[]interface{}", ""
		}
		typeName, pkgName := getTypeName(iType[0])
		return fmt.Sprintf("[]%s", typeName), pkgName
	default:
		return t.Name(), ""
	}
}
