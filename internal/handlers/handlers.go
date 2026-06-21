package handlers

import (
	handlers "github.com/DevanshBhavsar3/raven/internal/handlers/conversation"
	"github.com/DevanshBhavsar3/raven/internal/services"
)

type Handlers struct {
	Conversation *handlers.ConversationHandler
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		Conversation: handlers.NewConversationHandler(services.Conversation),
	}
}
