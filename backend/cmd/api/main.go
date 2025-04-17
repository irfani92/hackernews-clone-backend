package main

import (
	"hackernews-clone-backend/internal/adapter/driven/hnapi"
	"hackernews-clone-backend/internal/adapter/driver/httpn"
	"hackernews-clone-backend/internal/core/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const defaultPageSize = 20
const maxPageSize = 100

func main() {
	// Initialize Driven Adapters
	hnAPIAdapter := hnapi.NewHNAPIAdapter()

	// Initialize the Core Service, injecting the Driven Adapter
	itemService := service.NewItemService(hnAPIAdapter)

	// Initialize the Driver Adapter (HTTP Handler), injecting the Core Service
	itemHandler := httpn.NewItemHandler(itemService)

	// Set up the HTTP routing using Gin
	router := gin.Default()

	// Set up CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},         // Allow frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},  // Allowed HTTP methods
		AllowHeaders:     []string{"Content-Type", "Authorization"}, // Allowed headers
		AllowCredentials: true,                                      // Allow cookies or credentials if needed
	}))

	// Helper function to handle paginated list requests
	addPaginatedListRoute := func(relativePath string, handlerFunc gin.HandlerFunc) {
		router.GET(relativePath, func(c *gin.Context) {
			pageStr := c.DefaultQuery("page", "1")
			pageSizeStr := c.DefaultQuery("pageSize", strconv.Itoa(defaultPageSize))

			page, err := strconv.Atoi(pageStr)
			if err != nil || page < 1 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
				return
			}

			pageSize, err := strconv.Atoi(pageSizeStr)
			if err != nil || pageSize < 1 || pageSize > maxPageSize {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
				return
			}

			// Call the original handler with pagination parameters in the context
			c.Set("page", page)
			c.Set("pageSize", pageSize)
			handlerFunc(c)
		})
	}

	// Define routes
	router.GET("/items/:id", itemHandler.GetItem)
	addPaginatedListRoute("/topstories", itemHandler.ListTopStories)
	addPaginatedListRoute("/newstories", itemHandler.ListNewStories)
	addPaginatedListRoute("/askstories", itemHandler.ListAskStories)
	addPaginatedListRoute("/showstories", itemHandler.ListShowStories)
	addPaginatedListRoute("/jobstories", itemHandler.ListJobStories)

	// Start the HTTP server
	port := ":8080"
	log.Printf("Server listening on port %s\n", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
