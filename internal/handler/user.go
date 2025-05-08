package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// request, response schema
type (
	GetUsersResponse []GetUserResponse
	GetUserResponse  struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	}
)

// GET /api/v1/users
func (h *Handler) GetUsers(c echo.Context) error {
	users, err := h.repo.GetUsers(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	res := make(GetUsersResponse, len(users))
	for i, user := range users {
		userID, err := uuid.Parse(user.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
		}
		res[i] = GetUserResponse{
			ID:   userID,
			Name: user.Name,
		}
	}

	return c.JSON(http.StatusOK, res)
}
