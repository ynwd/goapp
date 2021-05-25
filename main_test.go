package main

import (
	"fmt"
	"testing"
)

var success = `<Agenda></Agenda>`

func TestGetName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "success",
			want: success,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetName()
			fmt.Println("g==>")
			fmt.Println("g==>", got)
			fmt.Println("w==>")
			fmt.Println("w==>", tt.want)

			if got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
