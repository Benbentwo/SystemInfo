package main

import (
	"github.com/Benbentwo/Windows10BootStrapper/app"
	"os"
)

func main() {
	if err := app.Run(nil); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
