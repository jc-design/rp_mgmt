package models

import (
	"fmt"
	"strings"
)

type Visibility int

const (
	Never    = 0
	Creation = Visibility(1) << iota
	Levelup
	Extended
	Other
)

func (vis *Visibility) ToString() string {
	var v []string
	*vis = *vis & (15)
	for i := 0; i < 4; i++ {
		if *vis&(1<<i) == 1 {
			v = append(v, "creation")
		} else if *vis&(1<<i) == 2 {
			v = append(v, "levelup")
		} else if *vis&(1<<i) == 4 {
			v = append(v, "extended")
		} else if *vis&(1<<i) == 8 {
			v = append(v, "other")
		}
	}
	if len(v) == 0 {
		v = append(v, "Never")
	}
	return strings.Join(v, "|")
}

func (vis *Visibility) FromString(value string) {
	*vis = Visibility(1)
	v := strings.Split(value, "|")
	for _, s := range v {
		if strings.ToLower(s) == "creation" && *vis > 0 {
			*vis = *vis | (1 << 0)
		} else if strings.ToLower(s) == "levelup" && *vis > 0 {
			*vis = *vis | (1 << 1)
		} else if strings.ToLower(s) == "extended" && *vis > 0 {
			*vis = *vis | (1 << 2)
		} else if strings.ToLower(s) == "other" && *vis > 0 {
			*vis = *vis | (1 << 3)
		} else {
			*vis = 0
		}
	}
}

// marshals a ElementTyp struct into a JSON string
func (vis *Visibility) MarshalJSON() ([]byte, error) {
	return []byte(addDoubleQuotes(vis.ToString())), nil
}

// unmarshals a JSON string into a ElementType struct
func (vis *Visibility) UnmarshalJSON(data []byte) error {
	vis.FromString(strings.Replace(string(data), "\"", "", -1))
	return nil
}

func addDoubleQuotes(s string) string {
	return fmt.Sprintf("\"%s\"", s)
}
