package models

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type Dice struct {
	DiceValue  int    `json:"dicevalue"`
	DiceCount  int    `json:"dicecount"`
	DiceMarkup int    `json:"dicemarkup"`
	Value      int    `json:"value"`
	Abr        string `json:"abr"`
}

func (d *Dice) SetValue(input ...any) {
	for _, val := range input {
		switch ass := val.(type) {
		case int:
			d.Value = ass
		case string:
			parsed, err := strconv.ParseInt(ass, 10, 64)
			if err == nil {
				d.Value = int(parsed)
			}
		case []int:
			if len(ass) >= 1 && ass[0] > 0 {
				d.DiceValue = ass[0]
			} else if len(ass) >= 2 && ass[1] > 0 {
				d.DiceCount = ass[1]
			} else if len(ass) >= 3 {
				d.DiceMarkup = ass[2]
			} else if len(ass) >= 4 {
				d.Value = ass[3]
			}
		default:
			fmt.Printf("could not work with input of type %T", ass)
		}
	}
}

func (d *Dice) GetInfo(key string) string {
	switch strings.ToLower(key) {
	case description:
		switch {
		case d.DiceMarkup == 0:
			return fmt.Sprintf("%d%s%d", d.DiceCount, d.Abr, d.DiceValue)
		case d.DiceMarkup < 0:
			return fmt.Sprintf("%d%s%d %d", d.DiceCount, d.Abr, d.DiceValue, d.DiceMarkup)
		case d.DiceMarkup > 0:
			return fmt.Sprintf("%d%s%d +%d", d.DiceCount, d.Abr, d.DiceValue, d.DiceMarkup)
		default:
			return ""
		}
	case identify:
		return fmt.Sprintf("%d%d%d", d.DiceValue, d.DiceCount, d.DiceMarkup)
	case value:
		return fmt.Sprintf("%d", d.Value)
	default:
		return ""
	}
}

func (d *Dice) Execute() (any, error) {
	//check conditions and return error if needed
	if !(d.DiceValue > 0 && d.DiceValue <= 100) {
		return nil, fmt.Errorf("dicevalue must be between 1 and 100 included")
	}
	if !(d.DiceMarkup >= -100 && d.DiceCount <= 100) {
		return nil, fmt.Errorf("dicemarku must be between -100 and 100 included")
	}

	//get random number as many times as DiceCount
	n := 0
	for range d.DiceCount {
		if r, err := rand.Int(rand.Reader, big.NewInt(int64(d.DiceValue))); err != nil {
			return nil, err
		} else {
			n += int(r.Int64()) + 1
		}
	}

	//add DiceMarkup
	n += d.DiceMarkup

	//return value
	return n, nil
}
