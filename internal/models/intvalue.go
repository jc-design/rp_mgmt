package models

import (
	"fmt"
	"strconv"
)

type IntValue struct {
	IntValue int `json:"intvalue"`
}

func (i *IntValue) SetValue(input ...any) {
	if len(input) > 0 {
		first_input := input[0]
		switch input := first_input.(type) {
		case int:
			i.IntValue = input
		case string:
			parsed, err := strconv.ParseInt(input, 10, 64)
			if err == nil {
				i.IntValue = int(parsed)
			}
		}
	}
}

func (i *IntValue) String() string {
	return fmt.Sprintf("%d", i.IntValue)
}

func (i *IntValue) InfosAsString() string {
	return ""
}

func (i *IntValue) Execute() {}
