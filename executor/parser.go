package executor

import (
	"io/ioutil"

	"github.com/zclconf/go-cty/cty"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	log "github.com/sirupsen/logrus"

	"github.com/wesovilabs/orion/dsl"
	"github.com/wesovilabs/orion/internal/errors"
)

// Parse will parse file content into valid config.
func (e *executor) Parse(path string) (*dsl.Feature, errors.Error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error(err.Error())
		return nil, errors.InvalidArguments("file cannot be read")
	}
	file, diagnostics := hclsyntax.ParseConfig(content, path, hcl.Pos{Line: 1, Column: 1, Byte: 0})
	if diagnostics != nil && diagnostics.HasErrors() {
		for i := range diagnostics.Errs() {
			log.Errorf("failed parsing file: '%s", diagnostics.Errs()[i].Error())
		}
		return nil, errors.Unexpected("file cannot be parsed")
	}
	out, decodeErr := e.dec.Decode(file.Body)
	if decodeErr != nil {
		return nil, decodeErr
	}
	out.SetPath(path)
	return out, nil
}

// ParseVariables will parse file content into map.
func ParseVariables(path string) (map[string]cty.Value, errors.Error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error(err.Error())
		return nil, errors.InvalidArguments("file cannot be read")
	}
	file, d := hclsyntax.ParseConfig(content, path, hcl.Pos{Line: 1, Column: 1, Byte: 0})
	if err := errors.EvalDiagnostics(d); err != nil {
		for i := range d.Errs() {
			log.Errorf("failed parsing file: '%s", d.Errs()[i].Error())
		}
		return nil, err
	}

	attributes, d := file.Body.JustAttributes()
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	output := make(map[string]cty.Value)
	for index := range attributes {
		attribute := attributes[index]
		val, d := attribute.Expr.Value(nil)
		if err := errors.EvalDiagnostics(d); err != nil {
			return nil, err
		}
		output[attribute.Name] = val
	}
	return output, nil
}
