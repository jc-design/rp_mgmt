package models

import (
	"fmt"
	"strings"
)

type Stringvalue struct {
	Stringvalue string `json:"stringvalue"`
}

func (s *Stringvalue) SetValue(input ...any) {
	for _, val := range input {
		switch ass := val.(type) {
		case string:
			s.Stringvalue = ass
		default:
			fmt.Printf("could not work with input of type %T", ass)
		}
	}
}

func (s *Stringvalue) GetInfo(key string) string {
	switch strings.ToLower(key) {
	case Value:
		return s.Stringvalue
	default:
		return ""
	}
}

func (s Stringvalue) Execute() (any, error) {
	return nil, nil
}
