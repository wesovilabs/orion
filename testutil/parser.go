package testutil

import (
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs/orion/internal/errors"
)

func GetStringAsBodyContent(text string, blockName string, labels []string) *hcl.BodyContent {
	parentSchema := &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{{
			Type:       blockName,
			LabelNames: labels,
		}},
	}
	file, err := ioutil.TempFile("", "orion.*.hcl")
	if err != nil {
		log.Errorf("error creating temporary file: %s", err)
		return nil
	}
	defer os.Remove(file.Name())

	parser := hclparse.NewParser()
	f, d := parser.ParseHCL([]byte(text), file.Name())
	if err := errors.EvalDiagnostics(d); err != nil {
		log.Error(err)
		return nil
	}
	content, d := f.Body.Content(parentSchema)
	if err := errors.EvalDiagnostics(d); err != nil {
		log.Error(err)
		return nil
	}
	return content
}
