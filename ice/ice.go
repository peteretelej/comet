package ice

import (
	"fmt"
	"log"
	"os/exec"
)

// InitProject initiliazes a comet project in the current working directory
func InitProject() error {
	if err := InitAssets(); err != nil {
		if err == ErrAppExists {
			return nil
		}
		return err
	}
	if err := GetElectron(); err != nil {
		return err
	}
	fmt.Println("comet: project initialized successfully.\nLaunch with `comet start`")
	return nil

}

// Launch starts app the application
func Launch(verbose bool, staticDir, staticURL string) error {
	Verbose = verbose
	if verbose {
		log.Print("comet: launching electron")
	}

	go func() {
		if staticURL != "" {
			UpdateURL(staticURL)
			fmt.Printf("comet: serving app from url: %s\n", staticURL)
			return
		}
		listen := "localhost:8080"
		if err := Serve(listen, staticDir); err != nil {
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
