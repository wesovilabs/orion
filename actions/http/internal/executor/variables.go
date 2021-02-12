package executor

import (
	"math/big"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)

type Variables interface {
	SetStatusCode(code int)
	SetBody(body []byte)
	SetToContext(ctx *hcl.EvalContext)
	SetHeaders(headers map[string][]string)
}

type variables struct {
	statusCode int
	body       []byte
	headers    map[string][]string
}

func createVariables() Variables {
	return &variables{}
}

func (v *variables) SetStatusCode(code int) {
	v.statusCode = code
}

func (v *variables) SetBody(body []byte) {
	v.body = body
}

func (v *variables) SetHeaders(headers map[string][]string) {
	v.headers = headers
}

func (v *variables) SetToContext(ctx *hcl.EvalContext) {
	statusCode := big.NewFloat(float64(v.statusCode))
	body := ""
	if v.body != nil {
		body = string(v.body)
	}

	headers := map[string]cty.Value{}
	for headerName, headerValue := range v.headers {
		if len(headerValue) == 1 {
			headers[headerName] = cty.StringVal(headerValue[0])
			continue
		}
		items := make([]cty.Value, len(headerValue))
		for index := range headerValue {
			items[index] = cty.StringVal(headerValue[index])
		}
		headers[headerName] = cty.ListVal(items)
	}
	httpVars := cty.ObjectVal(map[string]cty.Value{
		"body":       cty.StringVal(body),
		"headers":    cty.ObjectVal(headers),
		"statusCode": cty.NumberVal(statusCode),
	})
	if rootVars, ok := ctx.Variables["_"]; ok {
		rootValueMap := rootVars.AsValueMap()
		rootValueMap["http"] = httpVars
		ctx.Variables["_"] = cty.ObjectVal(rootValueMap)
		return
	}
	ctx.Variables["_"] = cty.ObjectVal(map[string]cty.Value{
		"http": httpVars,
	})
}
