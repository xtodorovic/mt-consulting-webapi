package consulting

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type implConsultingWaitingListAPI struct {
}

func NewConsultingWaitingListApi() RequestsListAPI {
	return &implConsultingWaitingListAPI{}
}

func (o implConsultingWaitingListAPI) CreateWaitingListEntry(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implConsultingWaitingListAPI) DeleteWaitingListEntry(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implConsultingWaitingListAPI) GetRequestsListEntries(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implConsultingWaitingListAPI) GetWaitingListEntry(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implConsultingWaitingListAPI) UpdateWaitingListEntry(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
