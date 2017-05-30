package ice

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// ErrAppExists is returned in electron app exists (main.js/package.json exists)
var ErrAppExists = errors.New("electron app already exists in directory")

const defaultURL = "http://localhost:8080"

var (
	appDir = filepath.Join("electron", "resources", "app")
)

// AssetsExist checks if any of the required electron assets already exist
func AssetsExist() bool {
	requiredAssets := []string{"main.js", "package.json"}
	for _, v := range requiredAssets {
		if _, err := os.Stat(filepath.Join(appDir, v)); err == nil {
			return true
		}
	}
	return false
}

// InitAssets creates static files required by comet init in the directory specified
// Skips prepping if all required files exist, and overwrites all if any is missing
func InitAssets() error {
	if AssetsExist() {
		if Verbose {
			log.Print("comet app assets already exist, skipping initialization")
		}
		return nil
	}
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return fmt.Errorf("unable to create app dir %s: %v", appDir, err)
	}
	pkgj, err := json.MarshalIndent(pj, "", "	")
	if err != nil {
		return fmt.Errorf("unable to parse package.json: %v", err)
	}
	required := map[string]string{
		"main.js":      fmt.Sprintf(mainJS, defaultURL),
		"package.json": string(pkgj),
	}
	for k, v := range required {
		fn := filepath.Join(appDir, k)
		err := ioutil.WriteFile(fn, []byte(v), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateURL replaces the url being served by the app by the one specified
// noop if url passed is empty string
func UpdateURL(url string) error {
	if url == "" {
		return nil
	}
	path := filepath.Join(appDir, "main.js")
	if _, err := os.Stat(path); err != nil {
		return err
	}
	s := fmt.Sprintf(mainJS, url)
	return ioutil.WriteFile(path, []byte(s), 0644)
}

const mainJS = `
const { app,BrowserWindow} = require('electron')

var win = null

app.on('ready', function(){
	win = new BrowserWindow({ width:800, height:600 })
	win.loadURL('%s')
})

app.on('window-all-closed', () => {
	app.quit()
})
`

type pkgJSON struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Main    string `json:"main"`
}

var pj = &pkgJSON{
	Name:    "@peteretelej/comet",
	Version: "1.0.0",
	Main:    "main.js",
}
