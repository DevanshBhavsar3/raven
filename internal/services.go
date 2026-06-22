package internal

import "github.com/DevanshBhavsar3/raven/internal/conversation"

type Services struct {
	Conversation *conversation.ConversationService
}

func NewServices(repositories *Repositories) *Services {
	return &Services{
		Conversation: conversation.NewConversationService(repositories.Conversation, repositories.Message),
	}
}
