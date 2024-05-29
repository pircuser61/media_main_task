package exchanges

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	name      string
	amount    int
	banknotes []int
	expect    []ExchangeRow
}

func (a ExchangeRow) equals(b ExchangeRow) bool {
	if len(a) != len(b) {
		return false
	}
	for i, x := range a {
		if x != b[i] {
			return false
		}
	}
	return true
}

func lessRows(i, j int, sl []ExchangeRow) bool {
	ll := len(sl[i])
	if ll != len(sl[j]) {
		return ll < len(sl[j])
	}
	// купюры в строках должны быть предварительно отсортированы
	for x := 0; x < ll; x++ {
		if sl[i][x] == sl[j][x] { // одинаковые пропускаем 500, 200 ~ 500, 100 /=> 200 > 100
			continue
		}
		return sl[i][x] < sl[j][x]
	}
	return false
}

func TestOne(t *testing.T) {
	cases := []testCase{
		{"нет размена", 11, []int{10, 50, 100}, nil},
		{"1 к 1", 10, []int{10}, []ExchangeRow{{10}}},
		{"какой то размен", 500, []int{1000, 50, 100, 200, 500},
			[]ExchangeRow{
				{500},

				{200, 200, 100},
				{200, 200, 50, 50},

				{200, 100, 100, 100},
				{200, 100, 100, 50, 50},
				{200, 100, 50, 50, 50, 50},
				{200, 50, 50, 50, 50, 50, 50},

				{100, 100, 100, 100, 100},
				{100, 100, 100, 100, 50, 50},
				{100, 100, 100, 50, 50, 50, 50},
				{100, 100, 50, 50, 50, 50, 50, 50},
				{100, 50, 50, 50, 50, 50, 50, 50, 50},

				{50, 50, 50, 50, 50, 50, 50, 50, 50, 50},
			},
		}}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			want := tc.expect
			got, err := GetExchages(tc.amount, tc.banknotes)
			require.NoError(t, err)
			// сравнить результаты наверное самое сложное в задании
			assert.Equal(t, len(want), len(got))
			// сортируем купюры во всех строках
			for i := 0; i < len(want); i++ {
				sort.Ints(want[i])
				sort.Ints(got[i])
			}
			lessWant := func(i, j int) bool {
				return lessRows(i, j, want)
			}
			lessGot := func(i, j int) bool {
				return lessRows(i, j, got)
			}
			// сортируем строки в ответах
			sort.Slice(want, lessWant)
			sort.Slice(got, lessGot)

			for i, w := range want {
				ok := w.equals(got[i])
				require.Truef(t, ok, "Строки должны совпадать", w, got[i])
			}

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
