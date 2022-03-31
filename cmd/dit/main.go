package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/alrs/dit"
)

var precise bool
var showVer bool
var sha string
var buildDate string

func init() {
	flag.BoolVar(&showVer, "v", false, "show version")
	flag.BoolVar(&precise, "p", false, "show desek precision")
	flag.Parse()
	if showVer {
		fmt.Printf("%s %s\n", sha, buildDate)
		os.Exit(0)
	}
}

func main() {
	now := time.Now()
	d := dit.TimeToDIT(now)
	if precise {
		fmt.Println(d)
	} else {
		fmt.Printf("%d.%d\n", d.Dec(), d.Decim())
	}
}
