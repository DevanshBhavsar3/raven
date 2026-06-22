package internal

import (
	"github.com/DevanshBhavsar3/raven/internal/conversation"
	"github.com/DevanshBhavsar3/raven/internal/conversation/message"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	Conversation *conversation.ConversationRepository
	Message      *message.MessageRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Conversation: conversation.NewConversationRepository(db),
		Message:      message.NewMessageRepository(db),
	}
}
