package httpn

import (
	"hackernews-clone-backend/internal/core/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	itemService *service.ItemService
}

func NewItemHandler(itemService *service.ItemService) *ItemHandler {
	return &ItemHandler{itemService: itemService}
}

func (h *ItemHandler) GetItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	item, err := h.itemService.GetItem(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch item"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) ListTopStories(c *gin.Context) {
	page, _ := c.Get("page")
	pageSize, _ := c.Get("pageSize")

	items, err := h.itemService.ListTopStories(c.Request.Context(), page.(int), pageSize.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch top stories"})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *ItemHandler) ListNewStories(c *gin.Context) {
	page, _ := c.Get("page")
	pageSize, _ := c.Get("pageSize")

	items, err := h.itemService.ListNewStories(c.Request.Context(), page.(int), pageSize.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch new stories"})
		return
	}
	c.JSON(http.StatusOK, items)
}

// Implement ListAskStories, ListShowStories, and ListJobStories similarly,
// passing the page and pageSize to the itemService methods.
func (h *ItemHandler) ListAskStories(c *gin.Context) {
	page, _ := c.Get("page")
	pageSize, _ := c.Get("pageSize")

	items, err := h.itemService.ListAskStories(c.Request.Context(), page.(int), pageSize.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch ask stories"})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *ItemHandler) ListShowStories(c *gin.Context) {
	page, _ := c.Get("page")
	pageSize, _ := c.Get("pageSize")

	items, err := h.itemService.ListShowStories(c.Request.Context(), page.(int), pageSize.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch show stories"})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *ItemHandler) ListJobStories(c *gin.Context) {
	page, _ := c.Get("page")
	pageSize, _ := c.Get("pageSize")

	items, err := h.itemService.ListJobStories(c.Request.Context(), page.(int), pageSize.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch job stories"})
		return
	}
	c.JSON(http.StatusOK, items)
}
