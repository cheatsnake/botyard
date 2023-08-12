package middlewares

import "botyard/internal/storage"

type Middlewares struct {
	store storage.Storage
}

func New(store storage.Storage) *Middlewares {
	return &Middlewares{
		store: store,
	}
}
