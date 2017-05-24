const { app,BrowserWindow} = require('electron')

var win = null

app.on('ready', function(){
	win = new BrowserWindow({ width:800, height:600 })
	win.loadURL('http://localhost:8080')
})

app.on('window-all-closed', () => {
	app.quit()
})

