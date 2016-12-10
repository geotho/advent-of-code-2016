package main

import (
	"reflect"
	"testing"
)

func Test_pointsBetween(t *testing.T) {
	type args struct {
		p1 point
		p2 point
	}
	tests := []struct {
		name string
		args args
		want []point
	}{
		{
			"same x",
			args{point{10, 10}, point{10, 13}},
			[]point{{10, 11}, {10, 12}, {10, 13}},
		},
		{
			"same y",
			args{point{10, 10}, point{13, 10}},
			[]point{{11, 10}, {12, 10}, {13, 10}},
		},
		{
			"same x backwards",
			args{point{10, 13}, point{10, 10}},
			[]point{{10, 12}, {10, 11}, {10, 10}},
		},
		{
			"same y backwards",
			args{point{13, 10}, point{10, 10}},
			[]point{{12, 10}, {11, 10}, {10, 10}},
		},
	}
	for _, tt := range tests {
		if got := pointsBetween(tt.args.p1, tt.args.p2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. pointsBetween() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
