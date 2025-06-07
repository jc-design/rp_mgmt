package models

type ValueElementer interface {
	ValueSetter
	Informer
	Executor
}

type ValueSetter interface {
	SetValue(...any) error
}

type Executor interface {
	Execute() (any, error)
}

type Informer interface {
	GetInfo(key string) (string, error)
}

type Cloner[C any] interface {
	Clone() C
}

func CloneAny[T Cloner[T]](c T) T {
	return c.Clone()
}

const (
	Description string = "description"
	Id          string = "id"
	Identify    string = "identify"
	Value       string = "value"
)
