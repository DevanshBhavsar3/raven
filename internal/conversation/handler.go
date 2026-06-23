package conversation

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DevanshBhavsar3/raven/internal/conversation/message"
	"github.com/DevanshBhavsar3/raven/internal/middlewares"
	"github.com/DevanshBhavsar3/raven/internal/utils"
	"github.com/go-chi/chi/v5"
)

type ConversationHandler struct {
	service *ConversationService
}

func NewConversationHandler(service *ConversationService) *ConversationHandler {
	return &ConversationHandler{
		service: service,
	}
}

func (h *ConversationHandler) CreateConversation(w http.ResponseWriter, r *http.Request) {
	payload := &CreateConversationRequest{}

	err := utils.BindAndValidate(r, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, ok := middlewares.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "user id not found in context", http.StatusUnauthorized)
		return
	}

	id, err := h.service.CreateConversation(r.Context(), userID, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	res := NewCreateConversationResponse(id)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ConversationHandler) GetAllConversations(w http.ResponseWriter, r *http.Request) {
	userID, ok := middlewares.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "user id not found in context", http.StatusUnauthorized)
		return
	}

	conversations, err := h.service.GetAllConversation(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := NewGetAllConversationsResponse(conversations)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ConversationHandler) GetConversationByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "conversationID")
	conversationID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "invalid conversation id", http.StatusBadRequest)
		return
	}

	userID, ok := middlewares.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "user id not found in context", http.StatusUnauthorized)
		return
	}

	conversation, err := h.service.GetConversationByID(r.Context(), userID, conversationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := NewGetConversationByIDResponse(conversation)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ConversationHandler) DeleteConversation(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "conversationID")
	conversationID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "invalid conversation id", http.StatusBadRequest)
		return
	}

	userID, ok := middlewares.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "user id not found in context", http.StatusUnauthorized)
		return
	}

	err = h.service.DeleteConversation(r.Context(), userID, conversationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ConversationHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "conversationID")
	conversationID, err := strconv.ParseInt(id, 10, 64)

	payload := &message.CreateMessageRequest{}

	err = utils.BindAndValidate(r, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, ok := middlewares.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "user id not found in context", http.StatusUnauthorized)
		return
	}
	_, err = h.service.GetConversationByID(r.Context(), userID, conversationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	messageID, err := h.service.CreateMessage(r.Context(), payload.Content, payload.Role, conversationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	res := message.NewCreateMessageResponse(messageID)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ConversationHandler) GetMessagesByConversationID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "conversationID")
	conversationID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "invalid conversation id", http.StatusBadRequest)
		return
	}

	_, err = h.service.GetConversationByID(r.Context(), "1", conversationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	messages, err := h.service.GetMessagesByConversationID(r.Context(), conversationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := message.NewGetMessagesByConversationIDResponse(messages)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
