package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/vimoj2/golibs"
)

func main() {
	log.Info("Loading application...")
	log.Warn("Pending...")
	log.Info(golibs.GetRandName())
}
