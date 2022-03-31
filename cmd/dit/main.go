package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/alrs/dit"
)

var precise bool

func init() {
	flag.BoolVar(&precise, "p", false, "show desek")
	flag.Parse()
}

func main() {
	now := time.Now()
	d := dit.TimeToDIT(now)
	if precise {
		fmt.Println(d)
	} else {
		fmt.Printf("%d:%d\n", d.Dec(), d.Decim())
	}
}
