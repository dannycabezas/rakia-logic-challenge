package main

import (
	"math/big"
	"strconv"
)

func numDecodings(s string) *big.Int {
	n := len(s)
	if n == 0 {
		return big.NewInt(0)
	}

	one_step_away := big.NewInt(1)
	two_steps_away := big.NewInt(0)
	current := big.NewInt(0)

	if s[n-1] != '0' {
		current.Set(one_step_away)
	} else {
		current.SetInt64(0)
	}

	two_steps_away.Set(one_step_away)
	one_step_away.Set(current)

	one_step_away = big.NewInt(1)
	two_steps_away = big.NewInt(0)

	for i := n - 1; i >= 0; i-- {
		current = new(big.Int)

		if s[i] == '0' {
			current.SetInt64(0)
		} else {

			current.Set(one_step_away)
		}

		if i+1 < n {
			twoDigitNum, err := strconv.Atoi(s[i : i+2])
			if err != nil {

			} else {
				if twoDigitNum >= 10 && twoDigitNum <= 26 {

					current.Add(current, two_steps_away)
				}
			}
		}

		two_steps_away.Set(one_step_away)
		one_step_away.Set(current)
	}

	return one_step_away
}
