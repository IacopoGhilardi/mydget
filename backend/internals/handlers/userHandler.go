package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/mydget-backend/internals/models"
	"github.com/iacopoghilardi/mydget-backend/internals/services"
	"github.com/iacopoghilardi/mydget-backend/internals/types/dto"
	"github.com/iacopoghilardi/mydget-backend/pkg/validation"
	"github.com/iacopoghilardi/mydget-backend/utils"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.userService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(
			"Internal Server Error",
			err.Error(),
		))
		return
	}
	c.JSON(http.StatusOK, utils.BuildSuccessResponse(users))
}

func (h *UserHandler) Create(c *gin.Context) {
	var dto dto.CreateUserDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"Bad Request",
			err.Error(),
		))
		return
	}

	if err := validation.ValidateCreateUserDto(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"Bad Request",
			err.Error(),
		))
		return
	}

	user, err := h.userService.Create(&dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(
			"Internal Server Error",
			err.Error(),
		))
		return
	}
	c.JSON(http.StatusOK, utils.BuildSuccessResponse(user))
}

func (h *UserHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"Bad Request",
			err.Error(),
		))
		return
	}
	user, err := h.userService.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(
			"Internal Server Error",
			err.Error(),
		))
		return
	}
	c.JSON(http.StatusOK, utils.BuildSuccessResponse(user))
}

func (h *UserHandler) Update(c *gin.Context) {
	var dto dto.UpdateUserDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"Bad Request",
			err.Error(),
		))
		return
	}
	user, err := h.userService.Update(models.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(
			"Internal Server Error",
			err.Error(),
		))
		return
	}
	c.JSON(http.StatusOK, utils.BuildSuccessResponse(user))
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"Bad Request",
			err.Error(),
		))
		return
	}
	err = h.userService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
