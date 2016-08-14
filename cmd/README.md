# Commands

## Notes

- use PersistPreRun to hook current Command and all its child Commands. https://github.com/spf13/cobra#prerun-or-postrun-hooks
- viper can lookup flags using `viper.BindPFlag`

## TODO

- [x] how to avoid same function call for every subcommand, ie: `Ayi git clone` does not trigger `Ayi git`, in `Ayi git` I 
read git related config into memory, I don't want to call it in `Ayi git clone` since there will also be 
`Ayi git sync`, `Ayi git rebase` etc. 
- solution is https://github.com/spf13/cobra#prerun-or-postrun-hooks