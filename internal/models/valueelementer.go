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
	description string = "description"
	id          string = "id"
	identify    string = "identify"
	value       string = "value"
)
