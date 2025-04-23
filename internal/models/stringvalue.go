package models

type StringValue struct {
	StringValue string `json:"stringvalue"`
}

func (i *StringValue) SetValue(input any) {
	switch input := input.(type) {
	case string:
		i.StringValue = input
	}
}

func (i *StringValue) ValueAsString() string {
	return i.StringValue
}

func (i StringValue) AdditionalValueAsString() string {
	return ""
}

func (i StringValue) Execute() {}
