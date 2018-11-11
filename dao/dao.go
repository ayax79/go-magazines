package dao

import (
	"github.com/google/uuid"

	"github.com/ayax79/go-magazines/model"
)

type (
	// MagazineDAO Interface for crud operations on magazine
	MagazineDAO interface {
		Get(uuid.UUID) (*model.Magazine, error)
		Put(*model.Magazine) error
	}
)
