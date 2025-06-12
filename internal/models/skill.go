package models

import (
	"fmt"
	"strconv"
	"strings"
)

var _ ValueElementer = (*Skill)(nil)
var _ Cloner[*Skill] = (*Skill)(nil)

type Skill struct {
	Skilltype       Skilltype `json:"skilltype"`
	DiceValue       int       `json:"dicevalue"`
	DiceCount       int       `json:"dicecount"`
	DiceMarkup      int       `json:"dicemarkup"`
	DiceBonusMarkup int       `json:"dicebonusmarkup"`
	Abr             string    `json:"abr"`
}

// ValueSetter interface
func (d *Skill) SetValue(input ...any) error {
	for _, val := range input {
		switch ass := val.(type) {
		case int:
			d.DiceMarkup = ass
			return nil
		case string:
			parsed, err := strconv.ParseInt(ass, 10, 64)
			if err == nil {
				d.DiceMarkup = int(parsed)
			}
			return err
		case []int:
			if len(ass) >= 1 && ass[0] != -999 {
				d.DiceValue = ass[0]
			}
			if len(ass) >= 2 && ass[1] != -999 {
				d.DiceCount = ass[1]
			}
			if len(ass) >= 3 && ass[2] != -999 {
				d.DiceMarkup = ass[2]
			}
			if len(ass) >= 4 && ass[3] != -999 {
				d.DiceBonusMarkup = ass[3]
			}
			return nil
		default:
			return fmt.Errorf("invalid type %T for SetValue @Dice. Value of input: %s", ass, ass)
		}
	}
	return fmt.Errorf("unkown error for SetValue @Dice")
}

// Informer interface
func (s *Skill) GetInfo(key string) (string, error) {
	switch strings.ToLower(key) {
	case Description:
		switch {
		case s.DiceMarkup == 0 && s.DiceBonusMarkup == 0:
			return fmt.Sprintf("%d%s%d", s.DiceCount, s.Abr, s.DiceValue), nil
		case s.DiceMarkup+s.DiceBonusMarkup < 0:
			return fmt.Sprintf("%d%s%d %d (%d)", s.DiceCount, s.Abr, s.DiceValue, s.DiceMarkup+s.DiceBonusMarkup, s.DiceBonusMarkup), nil
		case s.DiceMarkup > 0:
			return fmt.Sprintf("%d%s%d +%d (%d)", s.DiceCount, s.Abr, s.DiceValue, s.DiceMarkup+s.DiceBonusMarkup, s.DiceBonusMarkup), nil
		default:
			return "", fmt.Errorf("unkown error for GetValue @Typevalue")
		}
	case Identify:
		return fmt.Sprintf("%d%d%d", s.DiceValue, s.DiceCount, s.DiceMarkup), nil
	case Value:
		return fmt.Sprintf("%d", s.DiceMarkup+s.DiceBonusMarkup), nil
	default:
		return "", fmt.Errorf("wrong key (%s) for GetInfo @Skill", key)
	}
}

// Executor interface
func (s *Skill) Execute() (any, error) {
	return nil, nil
}

// Cloner interface
func (s *Skill) Clone() *Skill {
	return &Skill{
		Skilltype:       s.Skilltype,
		DiceValue:       s.DiceValue,
		DiceCount:       s.DiceCount,
		DiceMarkup:      s.DiceMarkup,
		DiceBonusMarkup: s.DiceBonusMarkup,
		Abr:             s.Abr,
	}
}

type Skilltype int

const (
	Innate int = iota
	Unlearned
	Learned
	Unknown
)

func (st *Skilltype) String() string {

	switch *st {
	case Skilltype(Innate):
		return "innate"
	case Skilltype(Unlearned):
		return "unlearned"
	case Skilltype(Learned):
		return "learned"
	default:
		return "unknown"
	}
}

func (st *Skilltype) FromString(value string) {
	var val Skilltype
	switch value {
	case "innate":
		val = Skilltype(Innate)
		st = &val
	case "unlearned":
		val = Skilltype(Unlearned)
		st = &val
	case "learned":
		val = Skilltype(Learned)
		st = &val
	default:
		val = Skilltype(Unknown)
		st = &val
	}
}

// marshals a ElementTyp struct into a JSON string
func (st *Skilltype) MarshalJSON() ([]byte, error) {
	return []byte(addDoubleQuotes(st.String())), nil
}

// unmarshals a JSON string into a ElementType struct
func (st *Skilltype) UnmarshalJSON(data []byte) error {
	st.FromString(strings.Replace(string(data), "\"", "", -1))
	return nil
}
