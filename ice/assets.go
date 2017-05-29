package ice

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ErrAppExists is returned in electron app exists (main.js/package.json exists)
var ErrAppExists = errors.New("electron app already exists in directory")

const defaultURL = "http://localhost:8080"

var appDir = filepath.Join("electron", "resources", "app")

// InitAssets creates static files required by comet init in the directory specified
// Skips prepping if all required files exist, and overwrites all if any is missing
func InitAssets() error {
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return fmt.Errorf("unable to create app dir %s: %v", appDir, err)
	}
	required := map[string]string{
		"main.js":      fmt.Sprintf(mainJS, defaultURL),
		"package.json": pkgJSON,
	}
	for k, v := range required {
		fn := filepath.Join(appDir, k)
		if _, err := os.Stat(fn); err == nil {
			return ErrAppExists
		}
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

const pkgJSON = `{
	"name": "@peteretelej/comet",
	"version": "1.0.0",
	"main": "main.js"
}
`
