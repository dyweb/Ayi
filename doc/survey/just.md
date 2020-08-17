# just

- [Github](https://github.com/casey/just)

## Overview

Syntax is very similar to `make`

Highlights

- can run in any folder, not the folder containing `justfile`
- can [pass arguments](https://github.com/casey/just#recipe-parameters)
- write recipes in [other language that has a shell](https://github.com/casey/just#writing-recipes-in-other-languages), i.e. set different shell (per receipt)
- [if](https://github.com/casey/just#if-statements), for and while

Make pitfalls

- `.PHONY`
- poor error message

## Internal

- a hand written LL(k) parser