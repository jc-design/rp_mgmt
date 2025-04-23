package models

import "fmt"

type ValueElementer interface {
	fmt.Stringer
	ValueSetter
	Informer
	Executor
}

type ValueSetter interface {
	SetValue(...any)
}

type Executor interface {
	Execute()
}

type Informer interface {
	InfosAsString() string
}
