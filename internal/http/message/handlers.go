package message

import "botyard/internal/storage"

type handlers struct {
	service *service
}

func Handlers(s storage.Storage) *handlers {
	return &handlers{
		service: &service{
			store: s,
		},
	}
}
