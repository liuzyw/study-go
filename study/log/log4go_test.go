package log


import (
	log "github.com/jeanphorn/log4go"
	"testing"
)

func Test_ddddd(t *testing.T) {
	log.LoadConfiguration("example.json")

	log.LOGGER("root").Info("category root Info test ...")
	log.LOGGER("root").Warn("category root Warn test message: %s", "new test msg")
	log.LOGGER("root").Debug("category root Debug test ...")
	log.LOGGER("root").Error("category root Error test ...")

	log.LOGGER("service").Info("category service Info test ...")
	log.LOGGER("service").Debug("category service Debug test ...")
	log.LOGGER("service").Warn("category service Warn test ...")

	log.LOGGER("warn").Info("category warn Info test ...")
	log.LOGGER("warn").Debug("category warn Debug test ...")
	log.LOGGER("warn").Warn("category warn Warn test ...")
	log.LOGGER("warn").Error("category warn Warn test ...")

	log.LOGGER("error").Warn("category error Warn test ...")
	log.LOGGER("error").Error("category error Warn test ...")


	log.LOGGER("not").Info("category not Info test ...")
	log.LOGGER("not").Debug("category not Debug test ...")
	log.LOGGER("not").Warn("category not Warn test ...")

	log.Close()
}

