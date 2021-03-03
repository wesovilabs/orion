package testutil

import (
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/wesovilabs/orion/internal/errors"
)

func GetBodyContent(path string, blockName string, labels []string) (*hcl.BodyContent, errors.Error) {
	parentSchema := &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{{
			Type:       blockName,
			LabelNames: labels,
		}},
	}
	parser := hclparse.NewParser()
	f, d := parser.ParseHCLFile(path)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	content, d := f.Body.Content(parentSchema)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	return content, nil
}

func GetAttributes(path string) (hcl.Attributes, errors.Error) {
	parser := hclparse.NewParser()
	f, d := parser.ParseHCLFile(path)
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	attributes, d := f.Body.JustAttributes()
	if err := errors.EvalDiagnostics(d); err != nil {
		return nil, err
	}
	return attributes, nil
}

func GetAttributesFromText(content string) (hcl.Attributes, error) {
	tmpFile, err := createTemporaryFileWith(content)
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())
	return GetAttributes(tmpFile.Name())
}

func MapStringAttributeToStringExpression(input map[string]*hcl.Attribute) map[string]hcl.Expression {
	out := make(map[string]hcl.Expression)
	for n, attr := range input {
		out[n] = attr.Expr
	}
	return out
}

func createTemporaryFileWith(content string) (*os.File, error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "pre-")
	if err != nil {
		return nil, err
	}
	text := []byte(content)
	if _, err = tmpFile.Write(text); err != nil {
		return nil, err
	}

	// Close the file
	if err := tmpFile.Close(); err != nil {
		return nil, err
	}
	return tmpFile, nil
}
