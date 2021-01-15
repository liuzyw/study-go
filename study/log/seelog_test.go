package log

import (
	log "github.com/cihub/seelog"
	"testing"
)

func init() {
	logger, err := log.LoggerFromConfigAsFile("seelog.xml")

	if err != nil {
		return
	}

	log.ReplaceLogger(logger)
}

func Test_df(t *testing.T) {
	defer log.Flush()
	log.Info("Hello from Seelog!")
	log.Warn("Warn")
	log.Error("Error")




}
