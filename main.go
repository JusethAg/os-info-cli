package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

// Custom type for splitting values from filter flag
type filter []string

// Method for formatting flag's value (Part of the flag.Value interface).
// Reference: https://golang.org/pkg/flag/#String
func (filters *filter) String() string {
	return fmt.Sprint(*filters)
}

// Method for setting the flag value (Part of the flag.Value interface).
// Reference: https://golang.org/pkg/flag/#Set
func (filters *filter) Set(value string) error {
	if len(*filters) > 0 {
		return errors.New("filters flag already set")
	}

	for _, filter := range strings.Split(value, ",") {
		*filters = append(*filters, filter)
	}

	return nil
}

func main() {
	getFlags()
}

func getFlags() {
	var all bool
	var filters filter

	flag.BoolVar(&all, "a", true, "Short access to 'all' flag")
	flag.BoolVar(&all, "all", true, "Show all info (CPU, Network and memory)")
	flag.Var(&filters, "f", "Short access to 'filter' flag")
	flag.Var(&filters, "filter", "cpu | net | mem")

	flag.Parse()

	fmt.Println("All flag value", all)
	fmt.Println("Filter flag value", filters)
}
