package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/peteretelej/comet/ice"
)

var (
	// startup flags
	verbose = flag.Bool("v", false, "verbose mode")
	static  = flag.String("static", "", "serve static directory (with index.html)")
	url     = flag.String("url", "", "serve a url on the desktop app (e.g. localhost:8080)")
)

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "start")
	}

	ice.Verbose = *verbose

	switch os.Args[1] {
	case "init":
		ice.Verbose = true
		fmt.Println("comet init: initializing your desktop app")
		if err := ice.InitProject(); err != nil {
			log.Fatalf("comet init: %v", err)
		}
		fmt.Println("comet app ready, launch with 'comet'")
	case "package":
		if err := packageApp(); err != nil {
			log.Fatalf("comet package: %v", err)
		}
	case "reset":
		if err := resetApp(); err != nil {
			log.Fatalf("comet reset: %v", err)
		}
	default:
		// start app
		err := ice.InitProject()
		if err != nil {
			log.Fatalf("comet: initialization failed: %v", err)
		}
		if err := ice.Launch(*static, *url); err != nil {
			fmt.Printf("comet start: %v\n", err)
			os.Exit(1)
		}
		return
	}
}

func packageApp() error {
	return fmt.Errorf("not yet implemented")

}

func resetApp() error {
	return fmt.Errorf("not yet implemented")
}
