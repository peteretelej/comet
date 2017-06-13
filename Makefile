build:
	CGO_ENABLED=0 GOOS=windows go build -o dist/comet_windows.exe main.go
	CGO_ENABLED=0 GOOS=linux go build -o dist/comet_linux main.go
	CGO_ENABLED=0 GOOS=darwin go build -o dist/comet_darwin main.go
	CGO_ENABLED=0 GOOS=freebsd go build -o dist/comet_freebsd main.go
