package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greenblat17/digital_spb/internal/entity"
)

func (h *Handler) ApplicantSignIn(c *gin.Context) {
	fmt.Println("sign in handler")
}

func (h *Handler) ApplicantSignUp(c *gin.Context) {//регистрация
	var input entity.Applicant
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
	}

	// id, err := h.services.Authorization.CreateAdmin(context.Background(), input)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	c.JSON(http.StatusOK, map[string]interface{}{
		//"id": id,
	})
}
