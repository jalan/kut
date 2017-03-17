# kut

```
$ kut -h
Usage: kut [OPTION] LIST

kut is a cut command for CSV (RFC 4180) input. kut reads from standard input
and writes to standard output.

LIST is a list of ranges separated by commas, using the same rules as cut(1).

Options:
  -d, --delimiter=DELIM    Use DELIM as field delimiters instead of commas.
  -h, --help               Show this help message.
```


## Why

cut(1) isn't sufficient for certain CSV files. There are some other tools that
solve this problem, but kut

 - has no dependencies
 - has a simple interface
 - only does one thing
 - is a statically linked executable


## How

Download a binary:

 - [linux x86](https://github.com/jalan/kut/releases/download/v1.0.0/kut_1.0.0_linux_x86)
 - [linux x86_64](https://github.com/jalan/kut/releases/download/v1.0.0/kut_1.0.0_linux_x86_64)

Or build one yourself:

```
go install github.com/jalan/kut/cmd/kut
```
