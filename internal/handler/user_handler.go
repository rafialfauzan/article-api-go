// internal/handler/user_handler.go
package handler

import (
	"golang-api/internal/model"
	"golang-api/internal/service"
	"golang-api/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.CreateUser(&user); err != nil {
		utils.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONSuccess(c, http.StatusCreated, "User created successfully", user)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		utils.JSONError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONSuccess(c, http.StatusOK, "Users retrieved successfully", users)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		utils.JSONError(c, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONSuccess(c, http.StatusOK, "User retrieved successfully", user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	updatedUser, err := h.service.UpdateUser(uint(id), &user)
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONSuccess(c, http.StatusOK, "User updated successfully", updatedUser)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.service.DeleteUser(uint(id)); err != nil {
		utils.JSONError(c, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONSuccess(c, http.StatusOK, "User deleted successfully", nil)
}
