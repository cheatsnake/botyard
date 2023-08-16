package handlers

import "botyard/internal/services"

type MessageHandlers struct {
	service *services.MessageService
}

func NewMessageHandlers(s *services.MessageService) *MessageHandlers {
	return &MessageHandlers{
		service: s,
	}
}
