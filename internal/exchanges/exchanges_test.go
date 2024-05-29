package exchanges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name      string
	amount    int
	banknotes []int
	expect    [][]RowItem
}

func TestOne(t *testing.T) {
	cases := []testCase{
		{"нет размена", 11, []int{10, 50, 100}, nil},
		{"1 к 1", 10, []int{10}, [][]RowItem{{{Val: 10, Count: 1}}}},
		{"какой то размен", 500, []int{1000, 50, 100, 200, 500},
			[][]RowItem{
				{{Val: 500, Count: 1}},

				{{Val: 200, Count: 2}, {Val: 100, Count: 1}},
				{{Val: 200, Count: 2}, {Val: 50, Count: 2}},

				{{Val: 200, Count: 1}, {Val: 100, Count: 3}},
				{{Val: 200, Count: 1}, {Val: 100, Count: 2}, {Val: 50, Count: 2}},
				{{Val: 200, Count: 1}, {Val: 100, Count: 1}, {Val: 50, Count: 4}},
				{{Val: 200, Count: 1}, {Val: 50, Count: 6}},

				{{Val: 100, Count: 5}},
				{{Val: 100, Count: 4}, {Val: 50, Count: 2}},
				{{Val: 100, Count: 3}, {Val: 50, Count: 4}},
				{{Val: 100, Count: 2}, {Val: 50, Count: 6}},
				{{Val: 100, Count: 1}, {Val: 50, Count: 8}},

				{{Val: 50, Count: 10}},
			},
		}}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := GetExchages(tc.amount, tc.banknotes)
			assert.NoError(t, err)
			assert.ElementsMatch(t, tc.expect, result)
		})
	}
}

func TestBigAmount(t *testing.T) {
	banknotes := []int{100, 200, 400, 600}
	amount := 100000 // потолок, на 1млн более 30с
	_, err := GetExchages(amount, banknotes)
	assert.NoError(t, err)

}

func TestBadAmout(t *testing.T) {
	_, err := GetExchages(0, []int{10, 50, 100})
	assert.ErrorIs(t, err, ErrBadAmount)
}

func TestEmptyBanknotes(t *testing.T) {
	_, err := GetExchages(10, nil)
	assert.ErrorIs(t, err, ErrEmptyBanknotes)
}

func TestBadBancknotes(t *testing.T) {
	_, err := GetExchages(10, []int{10, 50, 0})
	assert.ErrorIs(t, err, ErrBadBanknote)
}
