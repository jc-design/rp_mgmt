package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Intvalue struct {
	Intvalue int `json:"intvalue"`
}

func (i *Intvalue) SetValue(input ...any) {
	for _, val := range input {
		switch ass := val.(type) {
		case int:
			i.Intvalue = ass
		case string:
			parsed, err := strconv.ParseInt(ass, 10, 64)
			if err == nil {
				i.Intvalue = int(parsed)
			}
		default:
			fmt.Printf("could not work with input of type %T", ass)
		}
	}
}

func (i *Intvalue) GetInfo(key string) string {
	switch strings.ToLower(key) {
	case value:
		return fmt.Sprintf("%d", i.Intvalue)
	default:
		return ""
	}
}

func (i *Intvalue) Execute() (any, error) {
	return nil, nil
}
