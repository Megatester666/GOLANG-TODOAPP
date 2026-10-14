package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Megatester666/golang-todoapp/internal/core/logger"
	core_http_response "github.com/Megatester666/golang-todoapp/internal/core/transport/http/response"
	core_http_utils "github.com/Megatester666/golang-todoapp/internal/core/transport/http/utils"
)

type GetUsersResponse []UserDTOResponse

func (h *UsersHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'limit'/'offset' query params",
		)
		return
	}

	userDomains, err := h.usersService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'users' query params",
		)
		return
	}
	response := GetUsersResponse(usersDTOFromDomains(userDomains))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := core_http_utils.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("get limit: %w", err)
	}
	offset, err := core_http_utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("get offset: %w", err)
	}
	return limit, offset, nil
}
