package conversation

import "time"

type ConversationResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ########################################
// Create coversation response
// ########################################

type CreateConversationResponse struct {
	ID int64 `json:"id"`
}

func NewCreateConversationResponse(conversationID int64) CreateConversationResponse {
	return CreateConversationResponse{
		ID: conversationID,
	}
}

// ########################################
// Get all coversation response
// ########################################

type GetAllConversationsResponse []ConversationResponse

func NewGetAllConversationsResponse(conversations []Conversation) GetAllConversationsResponse {
	res := make(GetAllConversationsResponse, 0, len(conversations))

	for _, conversation := range conversations {
		res = append(res, ConversationResponse{
			ID:        conversation.ID,
			Name:      conversation.Name,
			UserID:    conversation.UserID,
			CreatedAt: conversation.CreatedAt,
			UpdatedAt: conversation.UpdatedAt,
		})
	}

	return res
}

// ########################################
// Get coversation by id response
// ########################################

type GetConversationByIDResponse ConversationResponse

func NewGetConversationByIDResponse(conversation *Conversation) GetConversationByIDResponse {
	return GetConversationByIDResponse{
		ID:        conversation.ID,
		Name:      conversation.Name,
		UserID:    conversation.UserID,
		CreatedAt: conversation.CreatedAt,
		UpdatedAt: conversation.UpdatedAt,
	}
}
