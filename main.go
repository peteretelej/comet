package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/peteretelej/comet/ice"
)

var (
	// subcommands
	initCommand  = flag.NewFlagSet("init", flag.ExitOnError)
	startCommand = flag.NewFlagSet("start", flag.ExitOnError)
	pkgCommand   = flag.NewFlagSet("package", flag.ExitOnError)
	resetCommand = flag.NewFlagSet("reset", flag.ExitOnError)

	// start subcommand flags
	verbose     = startCommand.Bool("v", false, "verbose mode")
	startStatic = startCommand.String("static", "", "serve static directory (with index.html)")
	startURL    = startCommand.String("url", "", "serve a url on the desktop app (e.g. localhost:8080)")
)

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "start")
	}
	switch os.Args[1] {
	case "init":
		if err := initProject(); err != nil {
			log.Fatalf("comet init: %v", err)
		}
	case "start":
		if err := startApp(); err != nil {
			log.Fatalf("comet start: %v", err)
		}
	case "package":
		if err := packageApp(); err != nil {
			log.Fatalf("comet package: %v", err)
		}
	case "reset":
		if err := resetApp(); err != nil {
			log.Fatalf("comet reset: %v", err)
		}
	default:
		err := initProject()
		if err != nil {
			log.Fatalf("comet: initialization failed: %v", err)
		}
		if err := startApp(); err != nil {
			log.Fatalf("comet: startup failed: %v", err)
		}
	}
}
func initProject() error {
	fmt.Println("comet: initializing project, this will take a few minutes..")
	if err := ice.InitAssets(); err != nil {
		if err == ice.ErrAppExists {
			fmt.Println("app already exists, skipping initialization.\n Launch with 'comet init'")
			return nil
		}
		return err
	}
	if err := ice.GetElectron(); err != nil {
		return err
	}
	fmt.Println("comet: project initialized successfully.\nLaunch with `comet start`")
	return nil
}

func startApp() error {
	var args []string
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}
	if err := startCommand.Parse(args); err != nil {
		return err
	}
	ice.Verbose = *verbose
	if *verbose {
		log.Print("comet: launching electron")
	}

	go func(url, staticDir string) {
		if url != "" {
			ice.UpdateURL(url)
			fmt.Printf("comet: serving app from url: %s\n", url)
			return
		}
		listen := "localhost:8080"
		if err := ice.Serve(listen, staticDir); err != nil {
			log.Printf("comet server crashed: %v", err)
		}
	}(*startURL, *startStatic)

	path, err := exec.LookPath("electron/electron")
	if err != nil {
		return fmt.Errorf("failed to find electron, did you run comet init?")
	}
	cmd := exec.Command(path, "electron")

	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("comet unable to launch electron: %s", out)
	}
	return nil
}

func packageApp() error {
	return fmt.Errorf("not yet implemented")

}

func resetApp() error {
	return fmt.Errorf("not yet implemented")
}
