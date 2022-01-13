package solution

import (
	"reflect"
	"testing"
)

func Test_mergeInBetween(t *testing.T) {
	type args struct {
		list1 *ListNode
		a     int
		b     int
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := mergeInBetween(tt.args.list1, tt.args.a, tt.args.b, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. mergeInBetween() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[0,1,2,3,4,5]
3
4
[1000000,1000001,1000002]
*/
