package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/TimothyYe/exchangerate"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

const defaultFrom = "USD"

var defaultTo = []string{"USD", "CNY", "SGD", "EUR", "GBP", "CAD", "AUD", "JPY"}
var defaultAPIKeyPath = filepath.Join(os.Getenv("HOME"), ".er")

func main() {
	checkAPIKey()
	apiKey, err := loadAPIKey()
	if err != nil || apiKey == "" {
		fmt.Println("failed to load API key")
		os.Exit(1)
	}

	from, amount, to := parseArgs(os.Args)
	result := exchangerate.Query(apiKey, from, amount, to)
	renderResult(from, amount, to, result)
}

func checkAPIKey() {
	if _, err := os.Stat(defaultAPIKeyPath); os.IsNotExist(err) {
		// key file does not exist
		fmt.Println("Cannot find API key file, please put your API key in file: ", defaultAPIKeyPath)
		os.Exit(1)
	}
}

func loadAPIKey() (string, error) {
	content, err := ioutil.ReadFile(defaultAPIKeyPath)
	if err != nil {
		return "", err
	}
	return strings.Replace(string(content), "\n", "", -1), nil
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
			displayLogo()
			displayHelp()
			os.Exit(0)
		} else if args[1] == "-v" || args[1] == "--version" {
			color.Cyan("Exchange Rate V%s", Version)
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
		amountArg := strings.Replace(args[2], ",", "", -1)
		s, err := strconv.ParseFloat(amountArg, 32)
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
		from = defaultFrom
		amount = 1
		to = defaultTo
	}

	return from, amount, to
}
