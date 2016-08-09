# Git 

This application is inspired by [hub](https://github.com/github/hub), but is much simpler. 
It can do the following:

- expand `clone` command 
- clone to your workspace regardless of your current directory
- add githooks

## TODO

- [x] support default http clone url in `remote.go`'s `NewFromURL`
- [ ] support default ssh clone url in `remote.go`'s `NewFromURL`
- [ ] handle non default ssh port
- [ ] detect workspace, config -> gopath -> askuser input and clone to workspace
- [ ] clone project by host order if using short url, and let user choose if there are same project on multiple hosts 