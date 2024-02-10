package handlers

import (
	"context"
	"net/http"

	"github.com/dwdarm/shortify/app/domain/services"
	"github.com/gin-gonic/gin"
)

type LinkHandler interface {
	LinkGet(c *gin.Context)
	LinkCreate(c *gin.Context)
}

type LinkHandlerImp struct {
	linkService services.LinkService
}

func NewLinkHandler(linkService services.LinkService) LinkHandler {
	return &LinkHandlerImp{
		linkService: linkService,
	}
}

func (h *LinkHandlerImp) LinkGet(c *gin.Context) {
	slug := c.Param("slug")
	link, err := h.linkService.GetLink(context.TODO(), slug)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})

		return
	}

	if link == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": link,
	})
}

type LinkCreateData struct {
	Slug string `form:"slug" json:"slug"`
	Href string `form:"href" json:"href" binding:"required"`
}

func (h *LinkHandlerImp) LinkCreate(c *gin.Context) {
	var form LinkCreateData
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	link, err := h.linkService.CreateLink(context.TODO(), form.Slug, form.Href)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": link,
	})
}
