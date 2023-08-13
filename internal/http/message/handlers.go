package message

type handlers struct {
	service *Service
}

func Handlers(s *Service) *handlers {
	return &handlers{
		service: s,
	}
}
