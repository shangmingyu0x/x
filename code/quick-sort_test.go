package code

import (
	"testing"
)

func Test_quickSort(t *testing.T) {
	var nilnums []int
	var zeronums = []int{}
	var equal = []int{1, 1, 2, 1, 1, 2}
	var normal = []int{9, 6, 4, 3, 1}
	type args struct {
		nums  *[]int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "nil",
			args: args{
				nums:  &nilnums,
				start: 0,
				end:   len(nilnums) - 1,
			},
		},
		{
			name: "zero",
			args: args{
				nums:  &zeronums,
				start: 0,
				end:   len(zeronums) - 1,
			},
		},
		{
			name: "equal",
			args: args{
				nums:  &equal,
				start: 0,
				end:   len(equal) - 1,
			},
		},
		{
			name: "normal",
			args: args{
				nums:  &normal,
				start: 0,
				end:   len(normal) - 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.nums, tt.args.start, tt.args.end)
			t.Logf("after: %v\n", tt.args.nums)
		})
	}
}
