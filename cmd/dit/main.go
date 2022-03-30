package main

import (
	"fmt"
	"time"

	"github.com/alrs/dit"
)

func main() {
	now := time.Now()
	d := dit.TimeToDIT(now)
	fmt.Println(d)
}
