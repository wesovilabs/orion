package decoder

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs/orion/helper"
	"github.com/wesovilabs/orion/internal/errors"
)

type Headers map[string]hcl.Expression

func (h Headers) buildHeaders(ctx *hcl.EvalContext) (map[string][]string, errors.Error) {
	headers := make(map[string][]string)
	for name := range h {
		val, d := h[name].Value(ctx)
		if err := errors.EvalDiagnostics(d); err != nil {
			return nil, err
		}
		if helper.IsSlice(val) {
			slice := val.AsValueSlice()
			headerValues := make([]string, len(slice))
			for i := range slice {
				headerValue, err := helper.ToString(slice[i])
				if err != nil {
					return nil, err
				}
				headerValues[i] = headerValue
			}
			headers[name] = headerValues
			continue
		}
		headerValue, err := helper.ToString(val)
		if err != nil {
			return nil, err
		}
		headers[name] = []string{headerValue}
	}
	return headers, nil
}
