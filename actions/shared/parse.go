package shared

import (
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

func CreatePos(line, col, b int) hcl.Pos {
	return hcl.Pos{
		Line:   line,
		Column: col,
		Byte:   b,
	}
}
