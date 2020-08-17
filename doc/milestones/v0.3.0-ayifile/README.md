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

### v0.2.5 Ayifile parser

Description



