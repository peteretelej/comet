# comet - Desktop Apps with Electron, Go, Bootstrap, Vuejs

Boostrap for desktop apps built with Electron and powered by Go. __WIP__

### Work In Progress
![stability-wip](https://img.shields.io/badge/stability-work_in_progress-lightgrey.svg)

This is __Work In Progress__: Not ready for use.


### Dev Requirements:

- Go
- Nodejs (npm)
	
### Basic Usage

Get the project
``` bash
git clone github.com/peteretelej/comet 
cd comet

# or via go get
go get -u github.com/peteretelej/comet
cd $GOPATH/src/github.com/peteretelej/comet
```


Initialize and launch
```
go build 
# compiles ./comet executable

./comet init
# initiliazes comet

./comet start
# launches app
```


### TODO

- [x] Define basic projects structure and working example
- [ ] ~~Spawn comet Go server from `main.js` __?__ ~~
- [x] Launch electron from Go
- [ ] Setup app templates for easier bootstrapping options
- [ ] __Packaging and distribution__

