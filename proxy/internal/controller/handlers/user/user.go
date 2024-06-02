package user

import (
	"GEO_API/proxy/internal/controller/responder"
	"GEO_API/proxy/internal/service/proxyService"
	"net/http"
)

type HandlerUser struct {
	s proxyService.User
	r responder.Responder
}

func New(service proxyService.User, responder responder.Responder) HandlerUser {
	return HandlerUser{service, responder}
}

// @Summary Profile
// @ID profile
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} models.User "get profile"
// @Failure 500
// @Security ApiKeyAuth
// @Router /user/profile [post]
func (h *HandlerUser) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	user, err := h.s.GetCurrentUser(r.Context())
	if err != nil {
		h.r.ErrorInternal(w, err)
		return
	}
	h.r.OutputJSON(w, responder.Response{
		Success: true,
		Message: "getting current user",
		Data:    user,
	})
}

// @Summary get list user
// @ID List
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} models.User "get list user"
// @Failure 500
// @Security ApiKeyAuth
// @Router /user/list [post]
func (h *HandlerUser) GetListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.s.GetListUsers()
	if err != nil {
		h.r.ErrorInternal(w, err)
		return
	}

	h.r.OutputJSON(w, responder.Response{
		Success: true,
		Message: "getting list users",
		Data:    users,
	})
}
