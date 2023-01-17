# go-app
Go App

## Steps to generate binaries.

### Windows
```shell
export VERSION=v0.2
env GOOS=windows GOARCH=amd64 go build -o go-app-windows-$VERSION.exe
```

### Linux
```shell
export VERSION=v0.2
env GOOS=windows GOARCH=amd64 go build -o go-app-linux-$VERSION
```

### Mac OS
```shell
export VERSION=v0.2
env GOOS=darwin GOARCH=amd64 go build -o go-app-darwin-$VERSION
```