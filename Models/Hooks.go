package Models

import "time"

type HookEvent int8

const (
	CreateHook HookEvent = 1 << iota
	UpdateHook           = 1 << iota
)

type ICreating interface {
	Creating()
}

type IUpdating interface {
	Updating()
}

/*
Hook Implementations
*/

func (m *BaseModel) Creating() {
	m.DateFields.CreatedAt = time.Now().UTC()
	m.DateFields.UpdatedAt = time.Now().UTC()
}

func (m *BaseModel) Updating() {
	m.DateFields.UpdatedAt = time.Now().UTC()
}
