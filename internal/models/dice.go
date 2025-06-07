package models

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

var _ ValueElementer = (*Dice)(nil)
var _ Cloner[*Dice] = (*Dice)(nil)

type Dice struct {
	DiceValue  int    `json:"dicevalue"`
	DiceCount  int    `json:"dicecount"`
	DiceMarkup int    `json:"dicemarkup"`
	Value      int    `json:"value"`
	Abr        string `json:"abr"`
}

// ValueSetter interface
func (d *Dice) SetValue(input ...any) error {
	for _, val := range input {
		switch ass := val.(type) {
		case int:
			d.Value = ass
			return nil
		case string:
			parsed, err := strconv.ParseInt(ass, 10, 64)
			if err == nil {
				d.Value = int(parsed)
			}
			return err
		case []int:
			if len(ass) >= 1 && ass[0] > 0 {
				d.DiceValue = ass[0]
			}
			if len(ass) >= 2 && ass[1] > 0 {
				d.DiceCount = ass[1]
			}
			if len(ass) >= 3 {
				d.DiceMarkup = ass[2]
			}
			if len(ass) >= 4 {
				d.Value = ass[3]
			}
			return nil
		default:
			return fmt.Errorf("invalid type %T for SetValue @Dice. Value of input: %s", ass, ass)
		}
	}
	return fmt.Errorf("unkown error for SetValue @Dice")
}

// Informer interface
func (d *Dice) GetInfo(key string) (string, error) {
	switch strings.ToLower(key) {
	case Description:
		switch {
		case d.DiceMarkup == 0:
			return fmt.Sprintf("%d%s%d", d.DiceCount, d.Abr, d.DiceValue), nil
		case d.DiceMarkup < 0:
			return fmt.Sprintf("%d%s%d %d", d.DiceCount, d.Abr, d.DiceValue, d.DiceMarkup), nil
		case d.DiceMarkup > 0:
			return fmt.Sprintf("%d%s%d +%d", d.DiceCount, d.Abr, d.DiceValue, d.DiceMarkup), nil
		default:
			return "", fmt.Errorf("unkown error for GetValue @Typevalue")
		}
	case Identify:
		return fmt.Sprintf("%d%d%d", d.DiceValue, d.DiceCount, d.DiceMarkup), nil
	case Value:
		return fmt.Sprintf("%d", d.Value), nil
	default:
		return "", fmt.Errorf("wrong key (%s) for GetInfo @Dice", key)
	}
}

// Executor interface
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

// Cloner interface
func (d *Dice) Clone() *Dice {
	return &Dice{
		DiceValue:  d.DiceValue,
		DiceCount:  d.DiceCount,
		DiceMarkup: d.DiceMarkup,
		Value:      d.Value,
		Abr:        d.Abr,
	}
}
