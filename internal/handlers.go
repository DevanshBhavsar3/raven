package internal

import (
	"github.com/DevanshBhavsar3/raven/internal/conversation"
)

type Handlers struct {
	Conversation *conversation.ConversationHandler
}

func NewHandlers(services *Services) *Handlers {
	return &Handlers{
		Conversation: conversation.NewConversationHandler(services.Conversation),
	}
}
