package log

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func init()  {

}

func Test_log1(t *testing.T) {
	log.Info("Hello from Seelog!,dasda, d%,",10,"dsada")
	log.Infof("Failed to send event %s to topic %s with key %d", "a", "b", 3)

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
		"n":10,
	}).Info("A group of walrus emerges from the ocean")

}