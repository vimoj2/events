package main

import (
	"fmt"
	"github.com/vimoj2/events/aggregate"
	"github.com/vimoj2/golibs"
)

func main() {
	fmt.Println("Hello world" + " " + golibs.GetRandName())
	fmt.Print("Hello world" + " " + aggregate.GetAgg())
}
