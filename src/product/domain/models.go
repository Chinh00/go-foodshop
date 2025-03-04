package domain

import "github.com/docker/distribution/uuid"

type (
	Product struct {
		Id    uuid.UUID `json:"id"`
		Name  string    `json:"name"`
		Price float64   `json:"price"`
	}
)
