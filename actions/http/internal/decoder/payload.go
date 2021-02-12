package decoder

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/url"

	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/helper"
	"github.com/wesovilabs/orion/internal/errors"
)

const (
	payloadJSON = "json"
	payloadXML  = "xml"
	payloadForm = "form"
	payloadRaw  = "raw"
)

type Payload struct {
	payloadType string
	data        hcl.Expression
}

func (p *Payload) SetType(t string) {
	p.payloadType = t
}

func (p *Payload) Type(ctx *hcl.EvalContext) string {
	return p.payloadType
}

func (p *Payload) SetData(expr hcl.Expression) {
	p.data = expr
}

func (p *Payload) Data(ctx *hcl.EvalContext) (string, errors.Error) {
	if p.payloadType == payloadRaw {
		return helper.GetExpressionValueAsString(ctx, p.data, "")
	}
	data, err := helper.GetExpressionValueAsInterface(ctx, p.data, nil)
	if err != nil {
		return "", err
	}
	pType := p.Type(ctx)
	switch pType {
	case payloadJSON:
		return p.toJSON(data)
	case payloadXML:
		return p.toXML(data)
	case payloadForm:
		return p.toForm(data)
	}
	return "", nil
}

func (p *Payload) toJSON(value interface{}) (string, errors.Error) {
	b, err := json.Marshal(value)
	if err != nil {
		return "", errors.Unexpected(err.Error())
	}
	return string(b), nil
}

func (p *Payload) toXML(value interface{}) (string, errors.Error) {
	b, err := xml.Marshal(value)
	if err != nil {
		return "", errors.Unexpected(err.Error())
	}
	return string(b), nil
}

func (p *Payload) toForm(value interface{}) (string, errors.Error) {
	form := url.Values{}
	switch dataMap := value.(type) {
	case map[string]interface{}:
		for k, v := range dataMap {
			form.Add(k, fmt.Sprintf("%v", v))
		}
	default:
		return "", errors.InvalidArguments("unsupported type for payload (form)")
	}
	return form.Encode(), nil
}
