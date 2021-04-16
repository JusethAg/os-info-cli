package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/JusethAg/os-info-cli/utils"
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

	for _, f := range strings.Split(value, ",") {
		*filters = append(*filters, f)
	}

	return nil
}

type flags struct {
	all     bool
	filters filter
}

func main() {
	GetFlagsFromCommandLine()

	netInfo := utils.GetNetworkInfo()

	fmt.Printf("Private IP: %v\n", netInfo.PrivateIp)
	fmt.Printf("Public IP: %v\n", netInfo.PublicIp)
}

func GetFlagsFromCommandLine() flags {
	var all bool
	var filters filter

	flag.BoolVar(&all, "a", true, "Short access to 'all' flag")
	flag.BoolVar(&all, "all", true, "Show all info (CPU, Network and memory)")
	flag.Var(&filters, "f", "Short access to 'filter' flag")
	flag.Var(&filters, "filter", "cpu | net | mem")

	flag.Parse()

	return flags{all, filters}
}
