package message

import "time"

type MessageResponse struct {
	ID             int64       `json:"id"`
	Content        string      `json:"content"`
	Role           MessageRole `json:"role"`
	ConversationID int64       `json:"conversationId"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAt"`
}

// ########################################
// Create message response
// ########################################

type CreateMessageResponse struct {
	ID int64 `json:"id"`
}

func NewCreateMessageResponse(messageID int64) CreateMessageResponse {
	return CreateMessageResponse{
		ID: messageID,
	}
}

// ########################################
// Get messages by conversation id response
// ########################################

type GetMessagesByConversationIDResponse []MessageResponse

func NewGetMessagesByConversationIDResponse(messages []Message) GetMessagesByConversationIDResponse {
	res := make(GetMessagesByConversationIDResponse, 0, len(messages))

	for _, message := range messages {
		res = append(res, MessageResponse{
			ID:             message.ID,
			Content:        message.Content,
			Role:           message.Role,
			ConversationID: message.ConversationID,
			CreatedAt:      message.CreatedAt,
			UpdatedAt:      message.UpdatedAt,
		})
	}

	return res
}
