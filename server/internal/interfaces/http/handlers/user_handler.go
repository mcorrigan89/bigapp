package handlers

import (
	"net/http"
	"net/mail"

	"github.com/google/uuid"
	"github.com/mcorrigan89/bigapp/server/internal/application"
	"github.com/mcorrigan89/bigapp/server/internal/application/queries"
	"github.com/mcorrigan89/bigapp/server/internal/interfaces/http/dto"
	"github.com/rs/zerolog"
)

type UserHandler struct {
	logger         *zerolog.Logger
	userAppService application.UserApplicationService
}

func NewUserHandler(logger *zerolog.Logger, userAppService application.UserApplicationService) *UserHandler {
	return &UserHandler{
		logger:         logger,
		userAppService: userAppService,
	}
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := r.PathValue("id")

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, "Failed to parse UUID", http.StatusInternalServerError)
		return
	}

	query := queries.UserByIDQuery{
		ID: userUUID,
	}

	user, err := h.userAppService.GetUserByID(ctx, query)
	if err != nil {
		http.Error(w, "Failed to get user by ID", http.StatusInternalServerError)
		return
	}

	userDto := dto.NewUserDtoFromEntity(user)

	userJson, err := userDto.ToJson()
	if err != nil {
		http.Error(w, "Failed to marshal user to JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(userJson)
}

func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userEmail := r.PathValue("email")

	_, err := mail.ParseAddress(userEmail)
	if err != nil {
		http.Error(w, "Failed to parse email", http.StatusInternalServerError)
		return
	}

	query := queries.UserByEmailQuery{
		Email: userEmail,
	}

	user, err := h.userAppService.GetUserByEmail(ctx, query)
	if err != nil {
		http.Error(w, "Failed to get user by email", http.StatusInternalServerError)
		return
	}

	userDto := dto.NewUserDtoFromEntity(user)

	userJson, err := userDto.ToJson()
	if err != nil {
		http.Error(w, "Failed to marshal user to JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(userJson)
}
