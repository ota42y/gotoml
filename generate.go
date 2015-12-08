package gotoml

import (
	"io"
	"fmt"
	"bytes"
	"sort"
	"reflect"
	"go/format"

	"github.com/BurntSushi/toml"
	"strings"
)

func Generate(input io.Reader, structName, pkgName string) ([]byte, error) {
	// read toml file
	var data map[string]interface{}
	if _, err := toml.DecodeReader(input, &data); err != nil {
		return nil, err
	}

	w := new(bytes.Buffer)
	generateHead(w, structName, pkgName)
	generateBody(w, data)
	fmt.Fprintf(w, "}")

	result, err := format.Source(w.Bytes())
	if err != nil {
		return nil, fmt.Errorf("go format error %s when %s formated", err.Error(), w.String())
	}

	return result, nil
}

func generateHead(w io.Writer, structName string, pkgName string) {
	fmt.Fprintf(w, "package %s\n\n", pkgName)
	fmt.Fprintf(w, "type %s struct {\n", structName)
}

func generateBody(w io.Writer, data map[string]interface{}) {
	mk := make([]string, len(data))
	i := 0
	for k, _ := range data {
		mk[i] = k
		i++
	}
	sort.Strings(mk)

	for _, key := range mk {
		keyTitle := strings.Title(key)
		typeName := reflect.TypeOf(key).Name()
		fmt.Fprintf(w, "%s %s `toml:\"%s\"`", keyTitle, typeName, key)
	}
}



