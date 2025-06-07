package models

import (
	"fmt"
	"strconv"
	"strings"
)

var _ ValueElementer = (*Intvalue)(nil)
var _ Cloner[*Intvalue] = (*Intvalue)(nil)

type Intvalue struct {
	Intvalue int `json:"intvalue"`
}

// ValueSetter interface
func (i *Intvalue) SetValue(input ...any) error {
	for _, val := range input {
		switch ass := val.(type) {
		case int:
			i.Intvalue = ass
			return nil
		case string:
			parsed, err := strconv.ParseInt(ass, 10, 64)
			if err == nil {
				i.Intvalue = int(parsed)
			}
			return err
		default:
			return fmt.Errorf("invalid type %T for SetValue @Intvalue. Value of input: %s", ass, ass)
		}
	}
	return fmt.Errorf("unkown error for SetValue @Intvalue")
}

// Informer interface
func (i *Intvalue) GetInfo(key string) (string, error) {
	switch strings.ToLower(key) {
	case Value:
		return fmt.Sprintf("%d", i.Intvalue), nil
	default:
		return "", fmt.Errorf("wrong key (%s) for GetInfo @Intvalue", key)
	}
}

// Executor interface
func (i *Intvalue) Execute() (any, error) {
	return nil, nil
}

// Cloner interface
func (i *Intvalue) Clone() *Intvalue {
	return &Intvalue{
		Intvalue: i.Intvalue,
	}
}
