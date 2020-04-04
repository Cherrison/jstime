package jstime

import (
	"reflect"
	"testing"
	"time"
)

func TestSetInterval(t *testing.T) {
	type args struct {
		d time.Duration
		f func()
	}
	var tests []struct {
		name string
		args args
		want *time.Ticker
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetInterval(tt.args.d, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetTimeout(t *testing.T) {
	type args struct {
		d time.Duration
		f func()
	}
	var tests []struct {
		name string
		args args
		want *time.Timer
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetTimeout(tt.args.d, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}