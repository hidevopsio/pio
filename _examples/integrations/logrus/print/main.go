package main

import (
	"github.com/hidevopsio/pio"
	_ "github.com/hidevopsio/pio/_examples/integrations/logrus"
)

func main() {
	pio.Print("This is an info message that will be printed to the logrus' printer")
}
