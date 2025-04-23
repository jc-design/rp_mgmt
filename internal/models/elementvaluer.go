package models

type ElementValuer interface {
	SetValue(any)
	ValueAsString() string
	AdditionalValueAsString() string
	Execute()
}
