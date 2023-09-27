package handler

import (
	_ "VikingsServer/docs"
	"VikingsServer/internal/app/repository"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

const (
	baseURL    = "api/v3"
	citiesHTML = "cities"

	cities           = baseURL + "/cities"
	addCityImage     = baseURL + "/cities/upload-image"
	hikes            = baseURL + "/hikes"
	users            = baseURL + "/users"
	DestinationHikes = baseURL + "/destination-hikes"
)

type Handler struct {
	Logger     *logrus.Logger
	Repository *repository.Repository
	Minio      *minio.Client
}

func NewHandler(l *logrus.Logger, r *repository.Repository, m *minio.Client) *Handler {
	return &Handler{
		Logger:     l,
		Repository: r,
		Minio:      m,
	}
}

func (h *Handler) RegisterHandler(router *gin.Engine) {
	router.GET(citiesHTML, h.CitiesHTML)
	router.POST(citiesHTML, h.DeleteCityHTML)

	router.GET(cities, h.CitiesList)
	router.POST(cities, h.AddCity)
	router.POST(addCityImage, h.AddImage)
	router.PUT(cities, h.UpdateCity)
	router.DELETE(cities, h.DeleteCity)

	router.GET(hikes, h.HikesList)
	router.POST(hikes, h.AddHike)
	router.DELETE(hikes, h.DeleteHike)
	router.PUT(hikes, h.UpdateHike)

	router.GET(users, h.UsersList)

	router.GET(DestinationHikes, h.DestinationHikesList)
	router.POST(DestinationHikes, h.AddDestinationToHike)

	registerStatic(router)
}

func registerStatic(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.LoadHTMLGlob("static/html/*")
	router.Static("/static", "./static")
	router.Static("/css", "./static")
	router.Static("/img", "./static")
}

// MARK: - Error handler

func (h *Handler) errorHandler(ctx *gin.Context, errorStatusCode int, err error) {
	h.Logger.Error(err.Error())
	ctx.JSON(errorStatusCode, gin.H{
		"status":      "error",
		"description": err.Error(),
	})
}

// MARK: - Success handler

func (h *Handler) successHandler(ctx *gin.Context, key string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		key:      data,
	})
}

func (h *Handler) successAddHandler(ctx *gin.Context, key string, data interface{}) {
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		key:      data,
	})
}
