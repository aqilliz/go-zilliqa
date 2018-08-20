Experimental golang implementation of Zilliqa.

[![Go Report Card](https://goreportcard.com/badge/github.com/aqilliz/go-zilliqa)](https://goreportcard.com/report/github.com/aqilliz/go-zilliqa)
[![Travis](https://travis-ci.org/aqilliz/go-zilliqa.svg?branch=master)](https://travis-ci.org/aqilliz/go-zilliqa)

## Go Zilliqa

Official golang implementation of Zilliqa.

## Building the source

Building gzil requires both a Go (version 1.7 or later) and a C compiler.
You can install them using your favourite package manager.
Once the dependencies are installed, run

    make gzil

or, to build the full suite of utilities:

    make all

## License

The go-zilliqa library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html), also
included in our repository in the `COPYING.LESSER` file.

The go-zilliqa binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included
in our repository in the `COPYING` file.
