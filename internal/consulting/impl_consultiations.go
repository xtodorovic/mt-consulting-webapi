package consulting

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
	"github.com/xtodorovic/mt-consulting-webapi/internal/db_service"
)

type implConsultationsAPI struct {
}

func NewConsultationsApi() ConsultationsAPI {
	return &implConsultationsAPI{}
}

func (o implConsultationsAPI) UpdateConsultation(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implConsultationsAPI) DeleteConsultation(c *gin.Context) {
	value, exists := c.Get("db_service")
	if !exists {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service not found",
				"error":   "db_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Consultation])
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service context is not of type db_service.DbService",
				"error":   "cannot cast db_service context to db_service.DbService",
			})
		return
	}

	requestId := c.Param("requestId")
	err := db.DeleteDocument(c, requestId)

	switch err {
	case nil:
		c.AbortWithStatus(http.StatusNoContent)
	case db_service.ErrNotFound:
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Consultation request not found",
				"error":   err.Error(),
			},
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to delete consultation request from database",
				"error":   err.Error(),
			})
	}
}

func (o implConsultationsAPI) SubmitConsultingForm(c *gin.Context) {
	value, exists := c.Get("db_service")
	if !exists {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Consultation])
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}
	form := Consultation{}
	err := c.BindJSON(&form)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}
	if form.Id == "" {
		form.Id = uuid.New().String()
	}

	err = db.CreateDocument(c, form.Id, &form)

	switch err {
	case nil:
		c.JSON(
			http.StatusCreated,
			form,
		)
	case db_service.ErrConflict:
		c.JSON(
			http.StatusConflict,
			gin.H{
				"status":  "Conflict",
				"message": "Consultation already exists",
				"error":   err.Error(),
			},
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create consultation in database",
				"error":   err.Error(),
			},
		)
	}
}

func (o implConsultationsAPI) GetRequestsListEntries(c *gin.Context) {
	value, exists := c.Get("db_service")
	if !exists {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service not found",
				"error":   "db_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Consultation])
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service context is not of type db_service.DbService",
				"error":   "cannot cast db_service context to db_service.DbService",
			})
		return
	}

	entries, err := db.ListDocuments(c)

	switch err {
	case nil:
		c.JSON(http.StatusOK, entries)
	case db_service.ErrNotFound:
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "No consultation requests found",
				"error":   err.Error(),
			},
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to fetch consultation requests from database",
				"error":   err.Error(),
			},
		)
	}
}
