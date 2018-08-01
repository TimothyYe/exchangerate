package main

import (
	"fmt"

	"github.com/fatih/color"
)

const (
	Version = "1.0"

	logo = `
 _                              _             
|_     _ |_   _. ._   _   _    |_)  _. _|_  _ 
|_ >< (_ | | (_| | | (_| (/_   | \ (_|  |_ (/_
                      _|    
Exchange Rate V%s
`
	help = `
https://github.com/TimothyYe/exchangerate

NAME:
   Exchange Rate - A simple command-line tool to query exchange rate

USAGE:
   er [from currency] [amount] [to currency]

VERSION:
   %s

EXAMPLES:
  er USD               Query USD and show exchange rate for common used currencies.  
  er USD 40.98         Query USD with amount 40.98, and show the equal amount of other currencies.  
  er USD 12 CNY,JPY    Query USD with amount 12 and show the equal amount of specified currencies.  

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
`
)

func displayLogo() {
	color.Cyan(logo, Version)
}

func displayHelp() {
	fmt.Printf(help, Version)
}
