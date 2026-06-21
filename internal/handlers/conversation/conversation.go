package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "github.com/DevanshBhavsar3/raven/internal/models/conversation"
	services "github.com/DevanshBhavsar3/raven/internal/services/conversation"
	"github.com/DevanshBhavsar3/raven/internal/utils"
	"github.com/go-chi/chi/v5"
)

type ConversationHandler struct {
	service *services.ConversationService
}

func NewConversationHandler(service *services.ConversationService) *ConversationHandler {
	return &ConversationHandler{
		service: service,
	}
}

func (h *ConversationHandler) CreateConversation(w http.ResponseWriter, r *http.Request) {
	payload := &models.CreateConversationRequest{}

	err := utils.BindAndValidate(r, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: use actual user id from auth context
	userID := "1"

	id, err := h.service.CreateConversation(r.Context(), userID, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	res := models.NewCreateConversationResponse(id)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ConversationHandler) GetAllConversations(w http.ResponseWriter, r *http.Request) {
	// TODO: use actual user id from auth context
	conversations, err := h.service.GetAllConversation(r.Context(), "1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := models.NewGetAllConversationsResponse(conversations)
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

	// TODO: use actual user id from auth context
	conversation, err := h.service.GetConversationByID(r.Context(), "1", conversationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := models.NewGetConversationByIDResponse(conversation)
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

	// TODO: use actual user id from auth context
	err = h.service.DeleteConversation(r.Context(), "1", conversationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
