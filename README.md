# kut

[![Docs](https://godoc.org/github.com/jalan/kut?status.svg)](https://godoc.org/github.com/jalan/kut)
[![Tests](https://github.com/jalan/kut/workflows/tests/badge.svg)](https://github.com/jalan/kut/actions)
[![Report Card](https://goreportcard.com/badge/github.com/jalan/kut)](https://goreportcard.com/report/github.com/jalan/kut)

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

```
go install github.com/jalan/kut/cmd/kut
```
