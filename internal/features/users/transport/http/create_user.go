package users_transport_http

import (
	"net/http"

	core_logger "github.com/Megatester666/golang-todoapp/internal/core/logger"
)

type createUserRequest struct {
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

type createUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	log.Debug("invoke CreateUser handler")
	var request createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println("error...")
	}

	rw.WriteHeader(http.StatusOK)
}
