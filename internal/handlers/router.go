package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/greenblat17/digital_spb/internal/service"
)

type Handler struct {
	service *service.Services
}

func NewHandler(service *service.Services) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("api/v1")
	{
		applicant := api.Group("/applicant")
		{
			applicant.POST("sign-up", h.ApplicantSignUp)
			applicant.GET("sign-in", h.ApplicantSignIn)
		}
		student := api.Group("/student")
		{
			student.POST("sign-up", h.StudentSignUp)
			student.GET("sign-in", h.StudentSignIn)
		}
	}

	return router
}
