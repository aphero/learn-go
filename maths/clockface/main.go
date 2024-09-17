package main

import (
	"os"
	"time"

	"learn-go/maths"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
