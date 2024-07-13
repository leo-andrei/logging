package drivers

import (
	"fmt"

	"github.com/leo-andrei/logging/log"
)

type CLIDriver struct{}

func (d CLIDriver) Log(entry log.LogEntry) {
	fmt.Println(entry)
}
