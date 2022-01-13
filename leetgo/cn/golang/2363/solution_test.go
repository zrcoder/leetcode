package solution

import (
	"reflect"
	"testing"
)

func Test_mergeSimilarItems(t *testing.T) {
	type args struct {
		items1 [][]int
		items2 [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := mergeSimilarItems(tt.args.items1, tt.args.items2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. mergeSimilarItems() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[[1,1],[4,5],[3,8]]
[[3,1],[1,5]]
*/
