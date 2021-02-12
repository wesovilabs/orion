package common

import (
	ct "github.com/daviddengcn/go-colortext"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wesovilabs/orion/internal/errors"
)

func PrintError(cmd *cobra.Command, err errors.Error) {
	cmd.Print()
	ct.ChangeColor(ct.Red, false, ct.None, false)
	log.Error(err)
}
