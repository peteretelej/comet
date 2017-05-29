# comet - Desktop Apps with Electron, Go, Bootstrap, Vuejs

Boostrap for desktop apps built with Electron and powered by Go. __WIP__

### Work In Progress
![stability-wip](https://img.shields.io/badge/stability-work_in_progress-lightgrey.svg)

This is __Work In Progress__: Not ready for use.


### Basic Usage

Get comet
``` bash
go get -u github.com/peteretelej/comet
```

Initialize and launch
```
comet init
# initiliazes comet

comet 
# starts app (initializes if needed)
```

### Launch Static Directory Single Page App/ PWA as desktop app
Assuming the directory ~/myapphtml is a static website with an index.html

```
# in any directory
comet init
comet -static ~/myapphtml
```

### Launch Website/ Web App as Desktop app

Serve a publicly accessible url as desktop app
```
comet -url https://etelej.com

```

- Note: changing the start URL (loadURL) is permanent (i.e. affects next run of `comet`),
  the default start url is `http://localhost:8080`, ie revert with `comet start -url http://localhost:8080`


### Other commands

```
comet reset 
# resets the comet installation on the directory
```


## TODO

- [x] Define basic projects structure and working example
- [x] Launch electron from Go
- [x] Setup comet CLI subcommands & usage (init,start)
- [x] Support serving static website as desktop app
- [x] Support serving abitrary url as app
- [ ] Add reset/ refresh command
- [ ] Setup app templates for easier bootstrapping options
- [ ] __Packaging and distribution__

