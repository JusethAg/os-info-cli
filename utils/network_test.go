package utils

import (
	"reflect"
	"testing"
)

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
