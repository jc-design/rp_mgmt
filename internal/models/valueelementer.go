package models

type ValueElementer interface {
	ValueSetter
	Informer
	Executor
}

type ValueSetter interface {
	SetValue(...any)
}

type Executor interface {
	Execute() (any, error)
}

type Informer interface {
	GetInfo(key string) string
}

const (
	Description string = "description"
	Id          string = "id"
	Identify    string = "identify"
	Value       string = "value"
)
