package models

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

type Dice struct {
	DiceValue  int    `json:"dicevalue"`
	DiceCount  int    `json:"dicecount"`
	DiceMarkup int    `json:"dicemarkup"`
	Value      int    `json:"value"`
	Abr        string `json:"abr"`
}

func (i *Dice) SetValue(input any) {
	switch input.(type) {
	case int:
		i.Value = input.(int)
	case string:
		parsed, err := strconv.ParseInt(input.(string), 10, 64)
		if err == nil {
			i.Value = int(parsed)
		}
	}
}

func (i *Dice) ValueAsString() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Dice) AdditionalValueAsString() string {
	switch {
	case i.DiceMarkup == 0:
		return fmt.Sprintf("%d%s%d", i.DiceCount, i.Abr, i.DiceValue)
	case i.DiceMarkup < 0:
		return fmt.Sprintf("%d%s%d %d", i.DiceCount, i.Abr, i.DiceValue, i.DiceMarkup)
	case i.DiceMarkup > 0:
		return fmt.Sprintf("%d%s%d +%d", i.DiceCount, i.Abr, i.DiceValue, i.DiceMarkup)
	default:
		return ""
	}
}

func (i *Dice) Execute() {
	if !(i.DiceValue > 0 && i.DiceValue <= 100) {
		return
	}
	if !(i.DiceMarkup >= -100 && i.DiceCount <= 100) {
		return
	}

	//get random number as many times as DiceCount
	n := 0
	for range i.DiceCount {
		if r, err := rand.Int(rand.Reader, big.NewInt(int64(i.DiceValue))); err != nil {
			return
		} else {
			n += int(r.Int64()) + 1
		}
	}

	//add DiceMarkup
	n += i.DiceMarkup

	//store result in Value
	i.Value = n
}
