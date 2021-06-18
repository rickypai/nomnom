package gen

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/rickypai/nomnom/gen/templates"
	"golang.org/x/tools/imports"
)

func ValuesStructTemplate(enum Enum) ([]byte, error) {
	return runTemplateBytes(templates.ValuesStruct, enum)
}

func ConversionsTemplate(enum Enum) ([]byte, error) {
	return runTemplateBytes(templates.Conversions, enum)
}

func ConversionsTestTemplate(enum Enum) ([]byte, error) {
	return runTemplateBytes(templates.ConversionsTest, enum)
}

func NumericConversionsTemplate(enum Enum) ([]byte, error) {
	return runTemplateBytes(templates.NumericConversions, enum)
}

func NumericConversionsTestTemplate(enum Enum) ([]byte, error) {
	return runTemplateBytes(templates.NumericConversionsTest, enum)
}

func runTemplateBytes(rawTemplate []byte, enum Enum) ([]byte, error) {
	t, err := template.New("letter").Parse(string(rawTemplate))
	if err != nil {
		return nil, fmt.Errorf("parsing template: %w", err)
	}

	var b bytes.Buffer

	err = t.Execute(&b, enum)
	if err != nil {
		return nil, fmt.Errorf("executing template: %w", err)
	}

	return b.Bytes(), nil
}

func formatCode(pkgName string, importedPkgs []string, content []byte) ([]byte, error) {
	result := []byte("// Code generated by nomnom. DO NOT EDIT.\n\npackage " + pkgName)
	result = append(result, "\n"[0])
	result = append(result, []byte("import (")...)

	for _, importedPkg := range importedPkgs {
		result = append(result, []byte("\t\""+importedPkg+"\"")...)
	}

	result = append(result, ")"[0])
	result = append(result, "\n"[0])

	result = append(result, content...)

	formatted, err := imports.Process("", result, nil)
	if err != nil {
		return nil, fmt.Errorf("running goimports: %w", err)
	}

	return formatted, nil
}
