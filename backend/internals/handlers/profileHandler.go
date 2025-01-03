package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/mydget-backend/internals/services"
	"github.com/iacopoghilardi/mydget-backend/internals/types/dto"
	"github.com/iacopoghilardi/mydget-backend/utils"
)

type ProfileHandler struct {
	profileService *services.ProfileService
}

func NewProfileHandler(profileService *services.ProfileService) *ProfileHandler {
	return &ProfileHandler{profileService}
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"Bad Request",
			err.Error(),
		))
		return
	}

	profile, err := h.profileService.GetProfile(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.BuildErrorResponse(
			"Not Found",
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusOK, utils.BuildSuccessResponse(profile))
}

func (h *ProfileHandler) CreateProfile(c *gin.Context) {
	var dto dto.CreateProfileDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"Bad Request",
			err.Error(),
		))
		return
	}

	profile, err := h.profileService.Create(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(
			"Internal Server Error",
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusOK, utils.BuildSuccessResponse(profile))
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	var dto dto.UpdateProfileDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"Bad Request",
			err.Error(),
		))
		return
	}

	profile, err := h.profileService.UpdateProfile(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(
			"Internal Server Error",
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusOK, utils.BuildSuccessResponse(profile))
}
