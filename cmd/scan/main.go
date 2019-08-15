package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/vimoj2/events/pkg/agg"
)

func main()  {
	fmt.Println("Logging...")
	log.Info("Info level")
	log.Info(agg.GetAgg());
}
