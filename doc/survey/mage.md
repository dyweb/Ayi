# Mage

- [Website](https://magefile.org/)
- [Github](https://github.com/magefile/mage)

## Take Away

- use as a library
- cache

## Overview

Write go file with `//+build mage` flag. 

Highlight

- [import](https://magefile.org/importing/)
- [dependencies](https://magefile.org/dependencies/) write `mg.Deps(func1, func2)`, it's dynamic at runtime ...
- [zeroinstall](https://magefile.org/zeroinstall/) use it as a library w/o a global binary
  - this should be the goal for gommon linter as well, it's much easier for user to customize behavior and manage version in go code

## Internal

I suppose it triggers go compiler w/ flag and inject a main func to call a target name. Use `go/parser` to find exported func as target name.

They got a page for [how it works](https://magefile.org/howitworks/), it has `~/.magefile` cache.