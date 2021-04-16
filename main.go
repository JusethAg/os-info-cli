package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
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

	for _, f := range strings.Split(value, ",") {
		*filters = append(*filters, f)
	}

	return nil
}

type flags struct {
	all     bool
	filters filter
}

type networkInfo struct {
	privateIp string
	publicIp  string
}

func main() {
	GetFlagsFromCommandLine()

	GetNetworkInfo()
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

func GetNetworkInfo() networkInfo {
	privateIp := getPrivateIP()
	publicIp := getPublicIp()

	return networkInfo{privateIp, publicIp}
}

func getPrivateIP() string {
	conn, err := net.Dial("tcp", "1.1.1.1:80")

	if err != nil {
		log.Fatal(err)
		return ""
	}

	defer conn.Close()
	tcpAddress, err := net.ResolveTCPAddr(
		conn.LocalAddr().Network(),
		conn.LocalAddr().String(),
	)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	ip := tcpAddress.IP.String()

	return ip
}

func getPublicIp() string {
	resp, err := http.Get("http://checkip.amazonaws.com")

	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	bodyResp, err := io.ReadAll(resp.Body)

	if err != nil {
		return ""
	}

	ip := string(bodyResp)

	return ip
}
