package models

import (
	"fmt"
	"strconv"
)

type IntValue struct {
	IntValue int `json:"intvalue"`
}

func (i *IntValue) SetValue(input any) {
	switch input := input.(type) {
	case int:
		i.IntValue = input
	case string:
		parsed, err := strconv.ParseInt(input, 10, 64)
		if err == nil {
			i.IntValue = int(parsed)
		}
	}
}

func (i *IntValue) ValueAsString() string {
	return fmt.Sprintf("%d", i.IntValue)
}

func (i *IntValue) AdditionalValueAsString() string {
	return ""
}

func (i *IntValue) Execute() {}
