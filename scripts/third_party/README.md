# Third party scripts

NOTE: Please open issue if you find including some scripts violate copyright.
I will remove them ASAP.

## Bash completion

- https://github.com/scop/bash-completion
- [GPLV2](https://github.com/scop/bash-completion/blob/master/COPYING)
- 2.3

When using cobra generated completion script, bash-completion must be installed, since 
it is using functions like `__ltrim_colon_completions`. 

After test on Windows with git bash, only `bash_completion` file is needed.