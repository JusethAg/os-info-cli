package utils

import (
	"reflect"
	"testing"
)

var mockPrivateIp = "127.0.0.1"
var mockpublicIp = "8.8.8.8"

func TestGetNetworkInfo(t *testing.T) {
	tests := []struct {
		name             string
		mockGetPrivateIP func() string
		mockGetPublicIP  func() string
		want             networkInfo
	}{
		{
			"Get private and public IP addresses",
			func() string { return mockPrivateIp },
			func() string { return mockpublicIp },
			networkInfo{mockPrivateIp, mockPrivateIp},
		},
		{
			"Get private IP address",
			func() string { return mockPrivateIp },
			func() string { return "" },
			networkInfo{mockPrivateIp, ""},
		},
		{
			"Get public IP address",
			func() string { return "" },
			func() string { return mockpublicIp },
			networkInfo{"", mockpublicIp},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetNetworkInfo()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNetworkInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
