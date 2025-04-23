package models

type StringValue struct {
	StringValue string `json:"stringvalue"`
}

func (i *StringValue) SetValue(input ...any) {
	if len(input) > 0 {
		first_input := input[0]
		switch input := first_input.(type) {
		case string:
			i.StringValue = input
		}
	}
}

func (i *StringValue) String() string {
	return i.StringValue
}

func (i StringValue) InfosAsString() string {
	return ""
}

func (i StringValue) Execute() {}
