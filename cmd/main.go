package main

import (
	"fmt"

	"github.com/pircuser61/media_main_task/internal/exchanges"
)

func main() {

	amount := 500
	b := []int{1000, 50, 100, 200, 500}

	x, _ := exchanges.GetExchages(amount, b)

	for _, r := range x {
		fmt.Println(r)

	}

}
