// +build tools

package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/tebeka/go2xunit"
	_ "golang.org/x/tools/cmd/goimports"
	_ "github.com/go-delve/delve/cmd/dlv"
)
