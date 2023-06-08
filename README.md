# bfcheck

---

**bfcheck** is a CLI tool for checking for Anti-BF FNF mods.

Before playing a FNF mod I run this util so that I don't get triggered.

## Anti-BF

An Anti-BF FNF mod is a mod where the opponent calls Boyfriend a midget or kills him.

## Building & Installing

1. Install Go

2. Run:
   
   ```shell
   go install github.com/MatusOllah/bfcheck@latest
   ```

## Usage

To check / scan a FNF mod run:

```shell
bfcheck -p path/to/fnf/mod
```

### Useful flags

- ``-p`` - path to FNF mod

- ``-v`` - prints verbose output

- ``-c`` - prints in color

- ``-l`` - shows found lines

- ``--version`` - prints version and exits