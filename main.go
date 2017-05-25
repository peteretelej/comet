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
		initProject()
	case "start":
		startApp()
	}
}
func initProject() {
	var args []string
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}
	if err := initCommand.Parse(args); err != nil {
		log.Fatal(err)
	}
	fmt.Println("comet: initializing project, may take a few seconds..")
	if _, err := exec.LookPath("npm"); err != nil {
		log.Fatal("Failed to find `npm`. Did you install nodejs?")
	}
	if err := ice.InitAssets(); err != nil {
		log.Fatalf("Failed to setup comet environment: %v", err)
	}
	cmd := exec.Command("npm", "install")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("comet: Running %v\n", cmd.Args)
	if err := cmd.Run(); err != nil {
		log.Fatalf("comet: failed to initialize project :(")
	}
	fmt.Println("comet: project initialized successfully. Launch with `comet start`")
}

func startApp() {
	var args []string
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}
	if err := startCommand.Parse(args); err != nil {
		log.Fatal(err)
	}
	if err := os.Chdir(*dir); err != nil {
		log.Fatalf("comet start: failed to change into directory: %v", err)
	}
	ice.Verbose = *verbose
	go func() {
		listen := "localhost:8080"
		ice.Serve(listen, *webapp)

	}()
	electron := filepath.Join("node_modules", ".bin", "electron")
	if _, err := os.Stat(electron); err != nil {
		log.Fatal("Failed to find electron in directory. Did you run `comet init`?")
	}
	if *verbose {
		log.Print("comet: launching electron")
	}

	if out, err := exec.Command(electron, ".").CombinedOutput(); err != nil {
		log.Fatalf("comet unable to launch electron: %s", out)
	}
}
