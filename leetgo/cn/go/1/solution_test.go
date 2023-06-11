package _two_sum

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. twoSum() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
