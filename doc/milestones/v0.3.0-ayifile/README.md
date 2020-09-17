# v0.3.0 Ayifile

## TODO

- [ ] list survey, there are many replacements for make
  - `justfile`
    - it might also contain link to alternatives
  - tj make
  - something used by go buffalo
  - bazel, buck etc.
- [ ] list feature

## Overview

`make` is a good build tool, but not an easy to use command runner. I want a config file format that allows defining commands easily.

## Motivation

- build a simple cli using config file instead of programming language
  - arg and flag handling, bash is a bit complex w/ getopt, make does not support flag and args
  - show help message without manually updating a large chunk of text (like in `make`)
- runs on windows w/o using external shell, i.e. consistent behaviour regardless of shell, cmd, ps, windows terminal, mysys32 etc.
  - bundle cross platform (core)util (written in go) like `sed`, `rm` etc.
- reuse across projects
  - e.g. most of my go projects have a similar layout (`cmd`, `pkg` etc.) and set of rules (`install`, `fmt`) etc. it would be good if I can `import lang/go.*`
 
## Features

### v0.2.3 Update dependencies

Description

Ayi is still using very old versions of dyweb/gommon and dyweb/go.ice. We need to remove go.ice (period) and use `go mod replace` for gommon.

Components

- `core`
  - following dyweb/pm, we can restructure the code to avoid the project level pkg.go
- `util`
  - most file in util can be removed, and the config loading logic etc. can be shared w/ pm via gommon

### v0.2.4 Ayifile spec

Description

We need to figure out the specification for `Ayifile`. There are many command runner and build tools.
If possible, similar or even same syntax (DSL) can be used in RCL and BenchHub.

Some features I have in mind

- dependencies, like `install: test` in make file
- tag, allow running a subset of commands, similar to tag in bazel target

### v0.2.5 Ayifile parser

Description



