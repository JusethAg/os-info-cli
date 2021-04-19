package utils

import (
	"reflect"
	"testing"
)

// TODO: Update tests refactoring implementation with interfaces and structs
func TestGetNetworkInfo(t *testing.T) {
	tests := []struct {
		name string
		want NetworkInfo
	}{
		{
			"Get private and public IP addresses",
			NetworkInfo{getPrivateIP(), getPublicIp()},
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
