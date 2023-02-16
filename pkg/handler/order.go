package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getJSONOrderById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	order, err := h.OrderService.GetByIdFromCache(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *Handler) getOrderById(c *gin.Context) {
	tmpl, err := template.ParseFiles("templates/order.html")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("failed parse template: %w", err).Error())
		return
	}

	id := c.Param("id")
	if id == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	order, err := h.OrderService.GetByIdFromCache(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = tmpl.Execute(c.Writer, order)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
}
