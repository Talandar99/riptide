# 🌊 Riptide
cli tool for running scripts 

# why not symlink?
- to eliminate risc of running specific scripts by accident. For example imagine that you have script that modify recursivly all files in current directory
- less junk in $PATH 

# TODO / current state
- [x] running scripts from predefined path
- [x] autocompletion
- [x] makefile or installation script
- [x] config file
- [ ] support for ~ (home directory)
- [ ] support for script arguments

## how to install
```
make install
```
- make will create config file at `$HOME/.config/riptide`, and scripts file at `$HOME/my_scripts` with 2 examples
## how to remove
```
make uninstall
```
## How it works / How to use
- After running for the first time it create configuration file in .config/riptide/conf.toml and .config/riptide/scripts
- You can specify path in toml file
- run by typing `riptide $scriptname`. You can use `<TAB>` for autocompletion

