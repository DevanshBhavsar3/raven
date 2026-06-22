package conversation

import (
	"context"

	"github.com/DevanshBhavsar3/raven/internal/conversation/message"
)

type ConversationService struct {
	conversationRepo *ConversationRepository
	messageRepo      *message.MessageRepository
}

func NewConversationService(conversationRepo *ConversationRepository, messageRepo *message.MessageRepository) *ConversationService {
	return &ConversationService{
		conversationRepo: conversationRepo,
		messageRepo:      messageRepo,
	}
}

func (s *ConversationService) CreateConversation(ctx context.Context, userID string, payload *CreateConversationRequest) (int64, error) {
	return s.conversationRepo.CreateConversation(ctx, userID, payload.Name)
}

func (s *ConversationService) GetAllConversation(ctx context.Context, userID string) ([]Conversation, error) {
	return s.conversationRepo.GetAllConversations(ctx, userID)
}

func (s *ConversationService) GetConversationByID(ctx context.Context, userID string, conversationID int64) (*Conversation, error) {
	return s.conversationRepo.GetConversationByID(ctx, userID, conversationID)
}

func (s *ConversationService) DeleteConversation(ctx context.Context, userID string, conversationID int64) error {
	return s.conversationRepo.DeleteConversation(ctx, userID, conversationID)
}

func (s *ConversationService) CreateMessage(ctx context.Context, content string, role message.MessageRole, conversationID int64) (int64, error) {
	return s.messageRepo.CreateMessage(ctx, content, role, conversationID)
}

func (s *ConversationService) GetMessagesByConversationID(ctx context.Context, conversationID int64) ([]message.Message, error) {
	return s.messageRepo.GetMessagesByConversationID(ctx, conversationID)
}
