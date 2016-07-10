# TODO

This file used to track TODO list 

## at15

### Applications

- [x] `git` should have `git.hosts` config in `.ayi.yml` parsed by order, but golang map does not have order. 
(@gaocegege has met this before) -> use array instead
- [ ] `viper` does not support array of `map[string]interface{}`, need to use `spf13/cast` to convert, which is also 
used by `viper`. see [git/config_test.go](git/config_test.go) `TestReadConfig` for detail

### Commands

- [ ] `install` and `update` should detect package manager like `npm`, `composer`

### Util

- [ ] https://github.com/uber-go/zap is a faster log library (TODO: the webframework I plan to use may also has its own logging library)
- [x] use https://github.com/Sirupsen/logrus for logging and use the `color` package in logrus if possible
- [x] better shell execute, see `github/hub/cmd`, which use `github.com/kballard/go-shellquote` for split and join