package http

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs/orion/actions"
	"github.com/wesovilabs/orion/actions/http/internal/decoder"
	"github.com/wesovilabs/orion/actions/http/internal/executor"
	"github.com/wesovilabs/orion/context"
	"github.com/wesovilabs/orion/helper"
	"github.com/wesovilabs/orion/internal/errors"
	"github.com/zclconf/go-cty/cty"
)

type HTTP struct {
	*actions.Base
	method   string
	request  *decoder.Request
	response *decoder.Response
}

func (h *HTTP) SetMethod(method string) {
	h.method = strings.ToUpper(method)
}

func (h *HTTP) Method() string {
	return strings.ToUpper(h.method)
}

func (h *HTTP) SetRequest(req *decoder.Request) {
	h.request = req
}

func (h *HTTP) SetResponse(res *decoder.Response) {
	h.response = res
}

func (p *HTTP) Execute(ctx context.FeatureContext) errors.Error {
	evalCtx := ctx.EvalContext()
	var err errors.Error
	var baseURL, path string
	executor := new(executor.HTTP)
	if baseURL, err = p.request.BaseURL(evalCtx); err != nil {
		return err
	}
	if path, err = p.request.Path(evalCtx); err != nil {
		return err
	}
	executor.URL = fmt.Sprintf("%s%s", baseURL, path)
	executor.Method = p.Method()
	if executor.Headers, err = p.request.Headers(evalCtx); err != nil {
		return err
	}
	if executor.QueryParams, err = p.request.QueryParams(evalCtx); err != nil {
		return err
	}
	if executor.Connection, err = p.request.Connection(evalCtx); err != nil {
		return err
	}
	if executor.Cookies, err = p.request.Cookies(evalCtx); err != nil {
		return err
	}

	log.Debugf("[%s] It starts the execution", BlockHTTP)

	vars, err := executor.Execute()
	if err != nil {
		return err
	}
	vars.SetToContext(evalCtx)
	if err := helper.EvaluateExpressions(evalCtx, p.response.Values()); err != nil {
		return err
	}
	cleanVariables(evalCtx)
	return nil
}

func cleanVariables(evalCtx *hcl.EvalContext) {
	variables := evalCtx.Variables
	rootVars := variables["_"].AsValueMap()
	delete(rootVars, BlockHTTP)
	variables["_"] = cty.ObjectVal(rootVars)
	evalCtx.Variables = variables
}
