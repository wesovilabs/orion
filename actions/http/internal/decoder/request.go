package decoder

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/wesovilabs-tools/orion/actions/http/internal/executor"
	"github.com/wesovilabs-tools/orion/helper"
	"github.com/wesovilabs-tools/orion/internal/errors"
)

type Request struct {
	baseURL     hcl.Expression
	path        hcl.Expression
	headers     Headers
	queryParams map[string]hcl.Expression
	connection  *Connection
	payload     *Payload
	cookies     []*Cookie
}

func (req *Request) BaseURL(ctx *hcl.EvalContext) (string, errors.Error) {
	return helper.GetExpressionValueAsString(ctx, req.baseURL, "")
}

func (req *Request) SetBaseURL(expr hcl.Expression) {
	req.baseURL = expr
}

func (req *Request) Path(ctx *hcl.EvalContext) (string, errors.Error) {
	return helper.GetExpressionValueAsString(ctx, req.path, "/")
}

func (req *Request) SetPath(expr hcl.Expression) {
	req.path = expr
}

func (req *Request) AddQueryParams(qp map[string]hcl.Expression) {
	if req.queryParams == nil {
		req.queryParams = make(QueryParams)
	}
	for name, value := range qp {
		req.queryParams[name] = value
	}
}

func (req *Request) AddHeaders(headers map[string]hcl.Expression) {
	if req.headers == nil {
		req.headers = make(Headers)
	}
	for name, value := range headers {
		req.headers[name] = value
	}
}

func (req *Request) Cookies(ctx *hcl.EvalContext) ([]*executor.Cookie, errors.Error) {
	cookies := req.cookies
	if len(cookies) == 0 {
		return nil, nil
	}
	var err errors.Error
	output := make([]*executor.Cookie, len(cookies))
	for index := range cookies {
		input := cookies[index]
		cookie := new(executor.Cookie)
		if cookie.Value, err = input.Value(ctx); err != nil {
			return nil, err
		}
		if cookie.Name, err = input.Name(ctx); err != nil {
			return nil, err
		}
		if cookie.Domain, err = input.Domain(ctx); err != nil {
			return nil, err
		}
		if cookie.Path, err = input.Path(ctx); err != nil {
			return nil, err
		}
		output[index] = cookie
	}
	return output, nil
}

func (req *Request) Connection(ctx *hcl.EvalContext) (*executor.Connection, errors.Error) {
	if req.connection == nil {
		return nil, nil
	}
	conn := new(executor.Connection)
	var err errors.Error
	if conn.Timeout, err = req.connection.Timeout(ctx); err != nil {
		return nil, err
	}
	if conn.Proxy, err = req.connection.Proxy(ctx); err != nil {
		return nil, err
	}
	return conn, nil
}

func (req *Request) Headers(ctx *hcl.EvalContext) (map[string][]string, errors.Error) {
	headers := make(map[string][]string)
	for name := range req.headers {
		val, d := req.headers[name].Value(ctx)
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

func (req *Request) QueryParams(ctx *hcl.EvalContext) (map[string][]string, errors.Error) {
	out := make(map[string][]string)
	for name := range req.queryParams {
		expr := req.queryParams[name]

		qpValue, d := expr.Value(ctx)
		if err := errors.EvalDiagnostics(d); err != nil {
			return nil, err
		}
		if helper.IsSlice(qpValue) {
			slice := qpValue.AsValueSlice()
			queryParamValues := make([]string, len(slice))
			for i := range slice {
				headerValue, err := helper.ToString(slice[i])
				if err != nil {
					return nil, err
				}
				queryParamValues[i] = headerValue
			}
			out[name] = queryParamValues
			continue
		}

		value, err := helper.ToString(qpValue)
		if err != nil {
			return nil, err
		}
		out[name] = []string{value}
	}
	return out, nil
}
