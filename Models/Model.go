package Models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model interface {
	GetID() primitive.ObjectID
}

type DateFields struct {
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type BaseModel struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	DateFields `bson:",inline"`
}

func (m BaseModel) GetID() primitive.ObjectID {
	return m.ID
}

/*
Fix for a possible bug with Generics??
*/

type IDFields interface {
	SetID(id interface{})
}

func (m *BaseModel) SetID(id interface{}) {
	m.ID = id.(primitive.ObjectID)
}

func setID(m Model, id interface{}) {
	if ins, ok := m.(IDFields); ok {
		ins.SetID(id)
	}
}
