# comet - Desktop Apps with Electron, Go, Bootstrap, Vuejs

Boostrap for desktop apps built with Electron and powered by Go. __WIP__

### Work In Progress
![stability-wip](https://img.shields.io/badge/stability-work_in_progress-lightgrey.svg)

This is __Work In Progress__: Not ready for use.


### Dev Requirements:

- Go
- Nodejs (npm)
	
### Basic Usage

Get comet
``` bash
go get -u github.com/peteretelej/comet
```

Initialize and launch
```
comet init
# initiliazes comet

comet start
# launches app
```

### Launch Static Website/ Single Page App/ PWA as desktop app
Assuming the directory ~/myapphtml is a static website with an index.html

```
cd ~/myapphtml
comet init
comet start -webapp .
```

### TODO

- [x] Define basic projects structure and working example
- [x] Launch electron from Go
- [x] Setup comet CLI subcommands & usage (init,start)
- [x] Support serving static website as desktop app
- [ ] Setup app templates for easier bootstrapping options
- [ ] __Packaging and distribution__

