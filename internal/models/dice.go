package models

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

type Dice struct {
	DiceValue  int
	DiceCount  int
	DiceMarkup int
	Value      int
	Abr        string
}

func (e *Dice) String() string {
	switch {
	case e.DiceMarkup == 0:
		return fmt.Sprintf("%d%s%d", e.DiceCount, e.Abr, e.DiceValue)
	case e.DiceMarkup < 0:
		return fmt.Sprintf("%d%s%d %d", e.DiceCount, e.Abr, e.DiceValue, e.DiceMarkup)
	case e.DiceMarkup > 0:
		return fmt.Sprintf("%d%s%d +%d", e.DiceCount, e.Abr, e.DiceValue, e.DiceMarkup)
	default:
		return ""
	}
}

func RollDice(d *Dice) (int, error) {

	if !(d.DiceValue > 0 && d.DiceValue <= 100) {
		return 0, errors.New("DiceCount must be between than 1 and 100")
	}
	if !(d.DiceCount > 0 && d.DiceCount <= 20) {
		return 0, errors.New("DiceCount must be between than 1 and 20")
	}
	if !(d.DiceMarkup >= -100 && d.DiceCount <= 100) {
		return 0, errors.New("DiceMArkup must be between than -100 and 100")
	}

	//get random number as many times as DiceCount
	n := 0
	for range d.DiceCount {
		if r, err := rand.Int(rand.Reader, big.NewInt(int64(d.DiceValue))); err != nil {
			return 0, err
		} else {
			n += int(r.Int64()) + 1
		}
	}

	//add DiceMarkup
	n += d.DiceMarkup

	//store result in Value
	d.Value = n
	return n, nil
}
