// Package diagnostics content relative to the package
package errors

import (
	"github.com/hashicorp/hcl/v2"
	log "github.com/sirupsen/logrus"
)

// EvalDiagnostics return error or nil depending on the diagnostics.
func EvalDiagnostics(d hcl.Diagnostics) Error {
	if d != nil && d.HasErrors() {
		for index := range d.Errs() {
			log.Warn(d.Errs()[index].Error())
		}
		return IncorrectUsage(d.Error())
	}

	return nil
}
