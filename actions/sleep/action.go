package sleep

import (
	"time"

	"github.com/hashicorp/hcl/v2"
	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs/orion/actions"
	"github.com/wesovilabs/orion/context"
	"github.com/wesovilabs/orion/helper"
	"github.com/wesovilabs/orion/internal/errors"
)

var defSleepDuration = 1 * time.Second

type Sleep struct {
	*actions.Base
	duration hcl.Expression
}

func (s *Sleep) SetDuration(expr hcl.Expression) {
	s.duration = expr
}

func (s *Sleep) Duration(ctx *hcl.EvalContext) (time.Duration, errors.Error) {
	sleepDuration, err := helper.GetExpressionValueAsDuration(ctx, s.duration, &defSleepDuration)
	return *sleepDuration, err
}

// Execute function in charge of executing the plugin.
func (s *Sleep) Execute(ctx context.OrionContext) errors.Error {
	return actions.Execute(ctx, s.Base, func(ctx context.OrionContext) errors.Error {
		duration, err := s.Duration(ctx.EvalContext())
		if err != nil {
			return err
		}

		log.Tracef("sleeping for duration %v", duration)
		time.Sleep(duration)

		return nil
	})
}
