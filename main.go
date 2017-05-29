package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/peteretelej/comet/ice"
)

var (
	// subcommands
	initCommand  = flag.NewFlagSet("init", flag.ExitOnError)
	startCommand = flag.NewFlagSet("start", flag.ExitOnError)

	// start subcommand flags
	verbose = startCommand.Bool("v", false, "verbose mode")
	dir     = startCommand.String("dir", "./", "directory to start comet from, where comet init was run")
	webapp  = startCommand.String("webapp", "", "serve static web app from directory instead of comet server")
)

func main() {
	flag.Parse()
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "start")
	}
	switch os.Args[1] {
	case "init":
		if err := initProject(); err != nil {
			log.Fatalf("comet init: %v", err)
		}
	case "start":
		startApp()
	case "package":
		packageApp()
	}
}
func initProject() error {
	fmt.Println("comet: initializing project, please wait..")
	if err := ice.GetElectron(); err != nil {
		return err
	}
	appDir := filepath.Join("resources", "app")
	if err := ice.InitAssets(appDir); err != nil {
		return err
	}

	fmt.Println("comet: project initialized successfully. Launch with `comet start`")
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
	if err := os.Chdir(*dir); err != nil {
		return fmt.Errorf("comet start: failed to change into directory: %v", err)
	}
	ice.Verbose = *verbose
	go func() {
		listen := "localhost:8080"
		if err := ice.Serve(listen, *webapp); err != nil {
			log.Printf("comet server crashed: %v", err)
		}
	}()
	electron := filepath.Join("node_modules", ".bin", "electron")
	if _, err := os.Stat(electron); err != nil {
		return fmt.Errorf("Failed to find electron in directory. Did you run `comet init`?")
	}
	if *verbose {
		log.Print("comet: launching electron")
	}

	if out, err := exec.Command(electron, ".").CombinedOutput(); err != nil {
		return fmt.Errorf("comet unable to launch electron: %s", out)
	}
	return nil
}

func packageApp() {

}
