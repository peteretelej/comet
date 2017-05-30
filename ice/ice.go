package ice

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// InitProject initiliazes a comet project in the current working directory
func InitProject() error {
	if err := InitAssets(); err != nil {
		return err
	}
	if err := GetElectron(); err != nil {
		return err
	}
	if Verbose {
		log.Print("comet: project initialized successfully`")
	}
	return nil

}

// Launch starts app the application
func Launch(staticDir, staticURL string) error {
	if Verbose {
		log.Print("comet: launching electron")
	}
	if staticDir != "" {
		if _, err := os.Stat(staticDir); err != nil {
			return fmt.Errorf("static directory specified '%s' not found", staticDir)
		}
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

	path, err := exec.LookPath(filepath.Join("electron", "electron"))
	if err != nil {
		return fmt.Errorf("failed to find electron, did you run comet init?")
	}
	cmd := exec.Command(path, "electron")

	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("comet unable to launch electron: %s", out)
	}
	return nil
}
