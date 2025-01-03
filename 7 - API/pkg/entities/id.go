package pkg_entities

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(s string) (ID, error) {
	id, e := uuid.Parse(s)
	return ID(id), e
}
