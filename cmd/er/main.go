package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/TimothyYe/exchangerate"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

const defaultFromt = "USD"

var defaultTo = []string{"USD", "CNY", "EUR", "GBP", "CAD", "AUD", "JPY"}

func main() {
	from, amount, to := parseArgs(os.Args)
	result := exchangerate.Query(from, amount, to)
	renderResult(from, amount, to, result)
}

func renderResult(from string, amount float32, to []string, result map[string]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	// Set header array
	header := []string{from}

	if len(to) > 0 {
		for _, t := range to {
			if t != from {
				header = append(header, t)
			}
		}
	} else {
		for _, t := range defaultTo {
			if t != from {
				header = append(header, t)
			}
		}
	}

	table.SetHeader(header)
	row := []string{color.RedString(fmt.Sprintf("%.2f", amount))}

	if len(to) > 0 {
		for _, t := range to {
			if t != from {
				row = append(row, color.CyanString(result[t]))
			}
		}
	} else {
		for _, t := range defaultTo {
			if t != from {
				row = append(row, color.CyanString(result[t]))
			}
		}
	}

	table.Append(row)
	table.Render() // Send output
}

func parseArgs(args []string) (string, float32, []string) {
	var from string
	var amount float32
	var to []string

	if len(args) == 2 {
		if args[1] == "-h" || args[1] == "--help" {
			fmt.Println("TODO: display help info")
			os.Exit(0)
		} else {
			if len(args[1]) != 3 {
				fmt.Println("Please input a valid currency")
				os.Exit(1)
			} else {
				from = args[1]
				amount = 1
				to = defaultTo
			}
		}
	}

	if len(args) >= 3 {
		from = args[1]
		s, err := strconv.ParseFloat(args[2], 32)
		if err != nil {
			fmt.Println("Please input valid amount")
		}
		amount = float32(s)

		if len(args) == 3 {
			to = defaultTo
		} else if len(args) == 4 {
			to = strings.Split(args[3], ",")
		}
	}

	if len(args) == 1 {
		from = defaultFromt
		amount = 1
		to = defaultTo
	}

	return from, amount, to
}
