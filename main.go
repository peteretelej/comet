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
	switch os.Args[1] {
	case "init":
		fmt.Println("comet init: initializing your desktop app")
		if err := initProject(); err != nil {
			log.Fatalf("comet init: %v", err)
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
		// start app
		err := initProject()
		if err != nil {
			log.Fatalf("comet: initialization failed: %v", err)
		}
		if err := startApp(*verbose, *static, *url); err != nil {
			log.Fatalf("comet start: %v", err)
		}
		return
	}

}
func initProject() error {
	if err := ice.InitAssets(); err != nil {
		if err == ice.ErrAppExists {
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

func startApp(verbose bool, staticDir, staticURL string) error {
	ice.Verbose = verbose
	if verbose {
		log.Print("comet: launching electron")
	}

	go func() {
		if staticURL != "" {
			ice.UpdateURL(staticURL)
			fmt.Printf("comet: serving app from url: %s\n", staticURL)
			return
		}
		listen := "localhost:8080"
		if err := ice.Serve(listen, staticDir); err != nil {
			log.Printf("comet server crashed: %v", err)
		}
	}()

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
