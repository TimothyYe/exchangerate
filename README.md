```
 _                              _             
|_     _ |_   _. ._   _   _    |_)  _. _|_  _ 
|_ >< (_ | | (_| | | (_| (/_   | \ (_|  |_ (/_
                      _|    
```
[![Build Status][1]][2] [![Go Report Card][7]][8] [![GoDoc][9]][10]

[1]: https://travis-ci.org/TimothyYe/exchangerate.svg?branch=master
[2]: https://travis-ci.org/TimothyYe/exchangerate
[7]: https://goreportcard.com/badge/github.com/timothyye/exchangerate
[8]: https://goreportcard.com/report/github.com/timothyye/exchangerate
[9]: https://godoc.org/github.com/TimothyYe/exchangerate?status.svg
[10]: https://godoc.org/github.com/TimothyYe/exchangerate

A simple command-line tool to query exchange rate.

![](https://github.com/TimothyYe/exchangerate/blob/master/snapshots/er-demo.gif?raw=true)

## Installation

#### Homebrew

```bash
brew tap timothyye/tap
brew install timothyye/tap/exchangerate
```

#### Using Go

```bash
go get github.com/TimothyYe/exchangerate/cmd/er
```

#### Manual Installation

Download it from [releases](https://github.com/TimothyYe/exchangerate/releases) and extact it to /usr/bin or your PATH directory.

## Get API Key

* Visit [https://free.currencyconverterapi.com/free-api-key](https://free.currencyconverterapi.com/free-api-key) and get your free API key.
* Save your API key into `~/.er` file.

## Usage
```bash
% er -h

Exchange Rate V1.1
https://github.com/TimothyYe/exchangerate

NAME:
   Exchange Rate - A simple command-line tool to query exchange rate

USAGE:
   er [from currency] [amount] [to currency]

VERSION:
   1.1

EXAMPLES:
  er USD               Query USD and show exchange rate for common used currencies.  
  er USD 40.98         Query USD with amount 40.98, and show the equal amount of other currencies.  
  er USD 12 CNY,JPY    Query USD with amount 12 and show the equal amount of specified currencies.  

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Licence

[MIT License](https://github.com/TimothyYe/exchangerate/blob/master/LICENSE)
