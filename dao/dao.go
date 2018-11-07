package dao

import (
	"github.com/google/uuid"
)

type (
	// MagazineDAO Interface for crud operations on magazine
	MagazineDAO interface {
		Get(uuid.UUID) (Magazine, error)
		Put(Magazine) error
	}
)
