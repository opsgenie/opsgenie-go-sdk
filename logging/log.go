package logging

import (
	"fmt"
	"github.com/cihub/seelog"
)

var logger seelog.LoggerInterface

func init() {
	DisableLog()
}

// DisableLog disables all library log output.
func DisableLog() {
	logger = seelog.Disabled
}

func UseLogger(newLogger seelog.LoggerInterface) {
	logger = newLogger
	seelog.UseLogger(logger)
}

func Logger() seelog.LoggerInterface {
	return logger
}

func ConfigureLogger(testConfig []byte) {
	loggr, err := seelog.LoggerFromConfigAsBytes([]byte(testConfig))
	if err != nil {
		fmt.Println("error occured: " + err.Error())
	}
	logger = loggr
	seelog.UseLogger(loggr)
}

// Call this before app shutdown
func FlushLog() {
	logger.Flush()
}
