package oriontest

import (
	"io/ioutil"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func ParseHCL(path string, schema *hcl.BodySchema) hcl.Blocks {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	file, diagnostics := hclsyntax.ParseConfig(content, path, hcl.Pos{Line: 1, Column: 1, Byte: 0})
	if diagnostics != nil && diagnostics.HasErrors() {
		panic(diagnostics.Error())
	}
	bodyContent, diagnostics := file.Body.Content(schema)
	if diagnostics != nil && diagnostics.HasErrors() {
		panic(diagnostics.Error())
	}
	return bodyContent.Blocks
}
