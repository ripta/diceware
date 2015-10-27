
## Downloading a Precompiled Binary

Head over to the [releases page](https://github.com/ripta/diceware/releases) to just get a binary.

## Usage

Usage is straight-forward:

```
$ diceware
1320 4260 4530 547 2323 3905
clown mommy now bater fk lotte
```

and you should find `diceware help` to be rather helpful.

You can see a list of compiled word lists:

```
$ diceware list
The following word lists are available:
  diceware
  diceware8k
```

which can be used like such:

```
$ diceware -w diceware8k
3631 1697 374 3706 3152 5801
kiev defer arty kudzu horny scoop
```

## Building from Source

This application cannot be `go get`, because it does not come with a word list. If you want a binary, head over to the releases page on Github. The following instructions are available to those that want to build from source.

Get a copy of this repository, without compiling:

```
$ go get -d github.com/ripta/diceware
$ cd $GOPATH/github.com/ripta/diceware
```

Get and install `go-bindata`, which is required to package the word list into the binary:

```
$ go get -u github.com/jteeuwen/go-bindata/...
```

Download the diceware word list, and use `go-bindata` to generate a file named `bindata.go` in this repository:

```
$ curl -sL -o data/diceware.asc 'http://world.std.com/~reinhold/diceware.wordlist.asc'
$ curl -sL -o data/diceware8k.txt 'http://world.std.com/~reinhold/diceware8k.txt'
$ go-bindata data
```

Build the binary:

```
$ go build
```

And if `$GOPATH/bin` is in your `$PATH`, then `diceware` should just work.

## New Word Lists

All files ending in `.txt` and `.asc` inside `data/` are available. A `.txt` file is assumed to be a straight list of one word per line, while a `.asc` file is assumed to be a tab-separated `garbage\tword` with one entry on each line.

Place the file in `data/`, name it correctly, rerun `go-bindata`, and `go build`. The new binary should show the new file under the `list` command.
