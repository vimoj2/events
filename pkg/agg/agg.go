package agg

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func GetAgg() string {
	fmt.Printf("Aggregate")
	log.Info("Aggregate")
	return "Aggregate"
}
