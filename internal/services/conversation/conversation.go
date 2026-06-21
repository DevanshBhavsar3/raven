package services

import (
	"context"

	models "github.com/DevanshBhavsar3/raven/internal/models/conversation"
	repositories "github.com/DevanshBhavsar3/raven/internal/repositories/conversation"
)

type ConversationService struct {
	repo *repositories.ConversationRepository
}

func NewConversationService(repo *repositories.ConversationRepository) *ConversationService {
	return &ConversationService{
		repo: repo,
	}
}

func (s *ConversationService) CreateConversation(ctx context.Context, userID string, payload *models.CreateConversationRequest) (int64, error) {
	return s.repo.CreateConversation(ctx, userID, payload.Name)
}

func (s *ConversationService) GetAllConversation(ctx context.Context, userID string) ([]models.Conversation, error) {
	return s.repo.GetAllConversations(ctx, userID)
}

func (s *ConversationService) GetConversationByID(ctx context.Context, userID string, conversationID int64) (*models.Conversation, error) {
	return s.repo.GetConversationByID(ctx, userID, conversationID)
}

func (s *ConversationService) DeleteConversation(ctx context.Context, userID string, conversationID int64) error {
	return s.repo.DeleteConversation(ctx, userID, conversationID)
}
