package main

import (
	"fmt"
	"github.com/vimoj2/events/pkg"
	"github.com/vimoj2/golibs"
)

func main() {
	fmt.Println("Hello world" + " " + golibs.GetRandName())
	fmt.Print("Hello world" + " " + pkg.GetAgg())
}
