# Web Interface for Ayi 

Issues: #15 #47

## Usage 

NOTE: Use `go get github.com/at15/go.rice/rice` instead of `go get github.com/GeertJohan/go.rice/rice` 
to work on windows, the fork disabled `zip -A Ayi.exe`

Embed static assets

- go build -o Ayi 
- rice append -i github.com/dyweb/Ayi/app/web --exec Ayi
- mv Ayi $GOPATH/bin/Ayi

## Functionalities

- [x] Serve static file `Ayi web static`
- [ ] Desktop notifications
- [ ] Home page
- [ ] log viewer

## TODO

- [x] allow simple http server
- [ ] embed static assets #47
- [ ] embed for multiple packages