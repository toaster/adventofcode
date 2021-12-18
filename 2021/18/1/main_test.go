package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tryExplode(t *testing.T) {
	for name, tt := range map[string]struct {
		number  snailfishNumber
		wantNum snailfishNumber
	}{
		"1st": {
			number:  parseSnailfishNumber("[[[[[9,8],1],2],3],4]"),
			wantNum: parseSnailfishNumber("[[[[0,9],2],3],4]"),
		},
		"2nd": {
			number:  parseSnailfishNumber("[7,[6,[5,[4,[3,2]]]]]"),
			wantNum: parseSnailfishNumber("[7,[6,[5,[7,0]]]]"),
		},
		"3rd": {
			number:  parseSnailfishNumber("[[6,[5,[4,[3,2]]]],1]"),
			wantNum: parseSnailfishNumber("[[6,[5,[7,0]]],3]"),
		},
		"4th": {
			number:  parseSnailfishNumber("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"),
			wantNum: parseSnailfishNumber("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"),
		},
		"5th": {
			number:  parseSnailfishNumber("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"),
			wantNum: parseSnailfishNumber("[[3,[2,[8,0]]],[9,[5,[7,0]]]]"),
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := tryExplode(tt.number)
			if tt.wantNum != nil {
				assert.True(t, got)
				assert.Equal(t, tt.wantNum, tt.number)
			} else {
				assert.False(t, got)
			}
		})
	}
}
