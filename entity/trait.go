package entity

import (
	"github.com/google/uuid"
)

type trait struct {
	id uuid.UUID `json:"-"`
}

func (t *trait) GetId() uuid.UUID {
	return t.id
}

func (t *trait) SetId(id uuid.UUID) {
	t.id = id
}
