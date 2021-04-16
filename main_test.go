package main

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

func TestGetFlags(t *testing.T) {
	tests := []struct {
		name      string
		inputArgs []string
		want      flags
	}{
		{"No flags", []string{"./os-info-cli"}, flags{all: true}},
		{"Set 'a' flag", []string{"./os-info-cli", "-a"}, flags{all: true}},
		{"Set 'all' flag", []string{"./os-info-cli", "--all"}, flags{all: true}},
		{"Set 'f' flag with empty values", []string{"./os-info-cli", "-f="}, flags{true, filter{""}}},
		{"Set 'f' flag with one value", []string{"./os-info-cli", "-f=cpu"}, flags{true, filter{"cpu"}}},
		{"Set 'f' flag with several values", []string{"./os-info-cli", "-f=cpu,mem,net"}, flags{true, filter{"cpu", "mem", "net"}}},
		{"Set 'filter' flag with empty values", []string{"./os-info-cli", "--filter="}, flags{true, filter{""}}},
		{"Set 'filter' flag with one value", []string{"./os-info-cli", "--filter=cpu"}, flags{true, filter{"cpu"}}},
		{"Set 'filter' flag with several values", []string{"./os-info-cli", "--filter=cpu,mem,net"}, flags{true, filter{"cpu", "mem", "net"}}},
		{"Set 'a' and 'f' flag with empty values", []string{"./os-info-cli", "-a", "-f="}, flags{true, filter{""}}},
		{"Set 'a' and 'f' flag with several values", []string{"./os-info-cli", "-a", "-f=cpu,mem"}, flags{true, filter{"cpu", "mem"}}},
		{"Set 'all' and 'filter' flag with empty values", []string{"./os-info-cli", "-all", "--filter="}, flags{true, filter{""}}},
		{"Set 'all' and 'filter' flag with several values", []string{"./os-info-cli", "-all", "-filter=cpu,mem"}, flags{true, filter{"cpu", "mem"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Reset OS arguments for next test execution
			// Reference: https://golang.org/pkg/flag/#pkg-variables
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

			// Simulate command line arguments
			os.Args = tt.inputArgs

			got := GetFlagsFromCommandLine()

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("TestGetFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNetworkInfo(t *testing.T) {
	tests := []struct {
		name string
		want networkInfo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNetworkInfo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNetworkInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
