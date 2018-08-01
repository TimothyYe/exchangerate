```
 .____               _                                       .___          .          
 /      _  .-   ___  /        ___  , __     ___.   ___       /   \   ___  _/_     ___ 
 |__.    \,'  .'   ` |,---.  /   ` |'  `. .'   ` .'   `      |__-'  /   `  |    .'   `
 |       /\   |      |'   ` |    | |    | |    | |----'      |  \  |    |  |    |----'
 /----/ /  \   `._.' /    | `.__/| /    |  `---| `.___,      /   \ `.__/|  \__/ `.___,
                                           \___/                                      
```
[![Release][3]][4] [![MIT licensed][5]][6] [![Build Status][1]][2] [![Go Report Card][7]][8] [![GoCover.io][11]][12] [![GoDoc][9]][10]

[1]: https://travis-ci.org/TimothyYe/exchange-rate.svg?branch=master
[2]: https://travis-ci.org/TimothyYe/exchange-rate
[3]: http://github-release-version.herokuapp.com/github/timothyye/exchange-rate/release.svg?style=flat
[4]: https://github.com/timothyye/exchange-rate/releases/latest
[5]: https://img.shields.io/dub/l/vibe-d.svg
[6]: LICENSE
[7]: https://goreportcard.com/badge/github.com/timothyye/exchange-rate
[8]: https://goreportcard.com/report/github.com/timothyye/exchange-rate
[9]: https://godoc.org/github.com/TimothyYe/exchange-rate?status.svg
[10]: https://godoc.org/github.com/TimothyYe/exchange-rate
[11]: https://img.shields.io/badge/gocover.io-81.8%25-green.svg
[12]: https://gocover.io/github.com/timothyye/exchange-rate

A simple command-line tool to query exchange rate.

![](https://github.com/TimothyYe/exchange-rate/blob/master/snapshots/er-demo.gif?raw=true)

## Installation

#### Homebrew

```bash
brew tap timothyye/tap
brew install timothyye/tap/exchange-rate
```

#### Using Go

```bash
go get github.com/TimothyYe/exchange-rate/cmd/er
```

#### Manual Installation

Download it from [releases](https://github.com/TimothyYe/exchange-rate/releases) and extact it to /usr/bin or your PATH directory.

## Usage
```bash
% er -h

Exchange Rate V0.1
https://github.com/TimothyYe/exchange-rate

NAME:
   Exchange Rate - A simple command-line tool to query exchange rate

USAGE:
   er [from currency] [amount] [to currency]

VERSION:
   0.1

EXAMPLES:
  er USD               Query USD and show exchange rate for common used currencies.  
  er USD 40.98         Query USD with amount 40.98, and show the equal amount of other currencies.  
  er USD 12 CNY,JPY    Query USD with amount 12 and show the equal amount of specified currencies.  

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Licence

[MIT License](https://github.com/TimothyYe/exchange-rate/blob/master/LICENSE)
