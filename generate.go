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
	// read toml file
	var data map[string]interface{}
	if _, err := toml.DecodeReader(input, &data); err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)
	packages := generateBody(body, structName, data)

	w := new(bytes.Buffer)
	generateHead(w, pkgName, packages)
	fmt.Fprintf(w, body.String())

	result, err := format.Source(w.Bytes())
	if err != nil {
		return nil, fmt.Errorf("go format error %s when %s formated", err.Error(), w.String())
	}

	return result, nil
}

func generateHead(w io.Writer, pkgName string, packages []string) {
	fmt.Fprintf(w, "package %s\n", pkgName)

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

func generateBody(w io.Writer, structName string, data map[string]interface{}) []string {
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
	pkgMap := make(map[string]bool)
	for _, key := range mk {
		keyTitle := strings.Title(key)
		typeName, pkgPath := getTypeName(data[key])

		// save package name
		pkgMap[pkgPath] = true

		fmt.Fprintf(w, "%s %s `toml:\"%s\"`\n", keyTitle, typeName, key)
	}
	fmt.Fprintf(w, "}")

	// return all package names
	packages := make([]string, 0)
	for k, _ := range pkgMap {
		packages = append(packages, k)
	}
	sort.Strings(packages)
	return packages
}

func getTypeName(i interface{}) (string, string) {
	t := reflect.TypeOf(i)
	pkgPath := t.PkgPath()
	if pkgPath == "" {
		name := t.Name()
		if name == "int" {
			name = "int64"
		}
		return name, ""
	}

	return fmt.Sprintf("%s.%s", t.PkgPath(), t.Name()), pkgPath
}
