package main

import (
	clockface "learn-go-with-tests/16-maths"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
