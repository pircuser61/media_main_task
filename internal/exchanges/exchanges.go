package exchanges

import (
	"errors"
	"sort"
)

var ErrBadAmount = errors.New("сумма должна быть положительным числом")
var ErrEmptyBanknotes = errors.New("не указаны номиналы")
var ErrBadBanknote = errors.New("номинал должен быть положительным числом")

type RowItem struct {
	Val   int
	Count int
}

func GetExchages(amount int, banknotes []int) ([][]RowItem, error) {
	if amount <= 0 {
		return nil, ErrBadAmount
	}

	if len(banknotes) < 1 {
		return nil, ErrEmptyBanknotes
	}

	sort.Ints(banknotes)
	if banknotes[0] <= 0 {
		return nil, ErrBadBanknote
	}
	banknotesLen := len(banknotes)
	counts := make([]int, banknotesLen)
	var result [][]RowItem

	var tryExchage func(int, int)
	tryExchage = func(amnt int, ix int) {
		count := 0 // сколько купюр данного номинала в размене, 0 - попытка размена только более мелкими
		amountLeft := amnt
		for ; amountLeft >= 0; count, amountLeft = count+1, amountLeft-banknotes[ix] {
			counts[ix] = count
			if amountLeft == 0 { // разменяли без остатка, добавляем в ответ
				row := make([]RowItem, 0, ix+1) // строка ответа будет не более ix+1
				for i := 0; i <= ix; i++ {
					if counts[i] > 0 {
						row = append(row, RowItem{Val: banknotes[i], Count: counts[i]})
					}
				}
				result = append(result, row)
				break
			}

			if ix+1 < banknotesLen {
				tryExchage(amountLeft, ix+1)
			}
		}

	}

	tryExchage(amount, 0)

	return result, nil
}
