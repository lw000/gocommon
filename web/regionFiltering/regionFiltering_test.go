package regionFiltering

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		regions []string
	}
	tests := []struct {
		name string
		args args
		want *RegionFiltering
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.regions...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
