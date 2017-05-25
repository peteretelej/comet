package ice

import (
	"io/ioutil"
	"os"
)

// InitAssets creates static files required by comet init
// Skips prepping if all required files exist, and overwrites all if any is missing
func InitAssets() error {
	required := map[string]string{
		"main.js":      mainJS,
		"package.json": pkgJSON,
	}
	ok := true
	for k := range required {
		if _, err := os.Stat(k); err != nil {
			ok = false
		}
	}
	if ok {
		return nil
	}
	for k, v := range required {
		err := ioutil.WriteFile(k, []byte(v), 0644)
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
