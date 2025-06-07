package models

import (
	"fmt"
	"strings"
)

var _ ValueElementer = (*Stringvalue)(nil)
var _ Cloner[*Stringvalue] = (*Stringvalue)(nil)

type Stringvalue struct {
	Stringvalue string `json:"stringvalue"`
}

// ValueSetter interface
func (s *Stringvalue) SetValue(input ...any) error {
	for _, val := range input {
		switch ass := val.(type) {
		case string:
			s.Stringvalue = ass
			return nil
		default:
			return fmt.Errorf("invalid type %T for SetValue @Stringvalue. Value of input: %s", ass, ass)
		}
	}
	return fmt.Errorf("unkown error for SetValue @Stringvalue")
}

// Informer interface
func (s *Stringvalue) GetInfo(key string) (string, error) {
	switch strings.ToLower(key) {
	case Value:
		return s.Stringvalue, nil
	default:
		return "", fmt.Errorf("wrong key (%s) for GetInfo @Typevalue", key)
	}
}

// Executor interface
func (s Stringvalue) Execute() (any, error) {
	return nil, nil
}

// Cloner interface
func (s *Stringvalue) Clone() *Stringvalue {
	return &Stringvalue{
		Stringvalue: s.Stringvalue,
	}
}
