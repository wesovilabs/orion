package logger

import (
	"fmt"

	ct "github.com/daviddengcn/go-colortext"
	log "github.com/sirupsen/logrus"
)

var levelColors = []ct.Color{
	ct.Red,     // PANIC
	ct.Red,     // FATAL
	ct.Red,     // ERROR
	ct.Magenta, //WARNING
	ct.Cyan,    // INFO
	ct.Green,   // DEBUG
	ct.Yellow,  // TRACE
}

type Formatter struct {
	TimestampFormat string
	ColorDisabled   bool
}

func (f *Formatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := fmt.Sprintf(entry.Time.Format(f.TimestampFormat))
	f.setColor(entry.Level)
	return []byte(fmt.Sprintf("%s %s\n", timestamp, entry.Message)), nil
}

func (f *Formatter) setColor(lvlIndex log.Level) {
	if !f.ColorDisabled {
		ct.ChangeColor(levelColors[lvlIndex], false, ct.None, false)
		return
	}
	ct.ResetColor()

}
