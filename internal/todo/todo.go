package todo

import (
	"time"
)

type Todo struct {
	ID       uint64    `json:"id"`
	Name     string    `json:"name"`
	Deadline time.Time `json:"date"`
}

type Repository interface {
}

type Service interface {
}
