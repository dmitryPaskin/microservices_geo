package auth

import (
	"GEO_API/proxy/internal/controller/responder"
	"GEO_API/proxy/internal/models"
	"GEO_API/proxy/internal/service/proxyService"
	"encoding/json"
	"net/http"
)

type HandlerAuth struct {
	s proxyService.Auth
	r responder.Responder
}

func New(service proxyService.Auth, responder responder.Responder) HandlerAuth {
	return HandlerAuth{service, responder}
}

// @Summary Register a user
// @ID SingUp
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.User true "User"
// @Success 201 "User registered successfully"
// @Failure 400 "Invalid request format"
// @Failure 500 "Response writer error on write"
// @Router /auth/register [post]
func (h *HandlerAuth) SingUpHandler(w http.ResponseWriter, r *http.Request) {
	var singUpUser models.User

	if err := json.NewDecoder(r.Body).Decode(&singUpUser); err != nil {
		h.r.ErrorBedRequest(w, err)
		return
	}

	if err := h.s.SingUpHandler(singUpUser); err != nil {
		h.r.ErrorInternal(w, err)
		return
	}

	h.r.StatusCreated(w)
}

// @Summary SingIn a user
// @ID SingIn
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.User true "User"
// @Success 200 "JWT token"
// @Failure 400 "Invalid request format"
// @Failure 500 "Response writer error on write"
// @Router /auth/login [post]
func (h *HandlerAuth) SingInHandler(w http.ResponseWriter, r *http.Request) {
	var singInUser models.User

	if err := json.NewDecoder(r.Body).Decode(&singInUser); err != nil {
		h.r.ErrorBedRequest(w, err)
		return
	}

	Token, err := h.s.SingInHandler(singInUser)
	if err != nil {
		h.r.ErrorInternal(w, err)
		return
	}

	h.r.OutputJSON(w, responder.Response{
		Success: true,
		Message: "Bearer: " + Token,
		Data:    nil,
	})
}
