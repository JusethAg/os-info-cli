package utils

import (
	"io"
	"log"
	"net"
	"net/http"
)

type NetworkInfo struct {
	privateIp string
	publicIp  string
}

func GetNetworkInfo() NetworkInfo {
	privateIp := getPrivateIP()
	publicIp := getPublicIp()

	return NetworkInfo{privateIp, publicIp}
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
