package internal

import (
	"github.com/DevanshBhavsar3/raven/internal/auth"
	"github.com/DevanshBhavsar3/raven/internal/config"
	"github.com/DevanshBhavsar3/raven/internal/conversation"
)

type Services struct {
	Auth         *auth.AuthService
	Conversation *conversation.ConversationService
}

func NewServices(cfg *config.ApplicationConfig, repositories *Repositories) *Services {
	return &Services{
		Auth:         auth.NewAuthService(cfg),
		Conversation: conversation.NewConversationService(repositories.Conversation, repositories.Message),
	}
}
