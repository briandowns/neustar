package main

// This is a WIP. Updates will be made to this and the README
// to reflect the API changes.

import (
	"fmt"
	"log"

	"github.com/briandowns/neustar"
)

func main() {
	conf, err := LoadConfig("./config.json")
	if err != nil {
		log.Fatalln(err)
	}
	monitor := neustar.NewMonitor(conf)
	fmt.Println(monitor.List())
	fmt.Println(monitor.Locations())
	fmt.Println(monitor.Get("4bbf505a660d11e49a049848e167c3b7"))
}
