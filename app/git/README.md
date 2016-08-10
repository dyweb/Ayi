# Git 

This application is inspired by [hub](https://github.com/github/hub), but is much simpler. 
It can do the following:

- expand `clone` command 
- clone to your workspace regardless of your current directory
- add githooks

## TODO

- [x] support default http clone url in `remote.go`'s `NewFromURL`
- [x] support default ssh clone url in `remote.go`'s `NewFromURL`
- [x] handle non default ssh port in `GetSSH`
- [x] detect workspace, config -> gopath -> ~~askuser input~~ current dir and clone to workspace
- [ ] clone project by host order if using short url, and let user choose if there are same project on multiple hosts 
- [ ] cleanup duplicated code `remote.go`, all the checking and error logging can be merged