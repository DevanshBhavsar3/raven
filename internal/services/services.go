package services

import (
	"github.com/DevanshBhavsar3/raven/internal/repositories"
	services "github.com/DevanshBhavsar3/raven/internal/services/conversation"
)

type Services struct {
	Conversation *services.ConversationService
}

func NewServices(repositories *repositories.Repositories) *Services {
	return &Services{
		Conversation: services.NewConversationService(repositories.Conversation),
	}
}
