package ice

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// InitAssets creates static files required by comet init in the directory specified
// Skips prepping if all required files exist, and overwrites all if any is missing
func InitAssets(dir string) error {
	fi, err := os.Stat(dir)
	if err != nil {
		return fmt.Errorf("specified directory %s does not exist", dir)
	}
	if !fi.IsDir() {
		return fmt.Errorf("cannot init assets to file speficied %s, must be a directory", dir)
	}
	required := map[string]string{
		"main.js":      mainJS,
		"package.json": pkgJSON,
	}
	for k, v := range required {
		fn := filepath.Join(dir, k)
		if _, err := os.Stat(fn); err == nil {
			continue
		}
		err := ioutil.WriteFile(fn, []byte(v), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

const mainJS = `
const { app,BrowserWindow} = require('electron')

var win = null

app.on('ready', function(){
	win = new BrowserWindow({ width:800, height:600 })
	win.loadURL('http://localhost:8080')
})

app.on('window-all-closed', () => {
	app.quit()
})
`

const pkgJSON = `{
  "name": "@peteretelej/comet",
  "version": "1.0.0",
  "description": "Build Desktop Apps with Electron, Golang, Bootstrap and Vuejs",
  "main": "main.js",
  "repository": {
    "type": "git",
    "url": "github.com/peteretelej/comet"
  },
  "keywords": [
    "comet",
    "electron",
    "golang",
    "vuejs",
    "bootstrap",
    "go"
  ],
  "author": "Peter Etelej",
  "license": "ISC",
  "dependencies": {
    "electron": "^1.6.8"
  },
  "devDependencies": {},
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  }
}
`
