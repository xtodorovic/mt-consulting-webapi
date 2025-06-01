package consulting_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/xtodorovic/mt-consulting-webapi/internal/consulting"
)

// MockDbService is a mock of db_service.DbService[Consultation]
type MockDbService struct {
	mock.Mock
}

func (m *MockDbService) FindDocument(c *gin.Context, id string) (consulting.Consultation, error) {
	args := m.Called(c, id)
	return args.Get(0).(consulting.Consultation), args.Error(1)
}

func (m *MockDbService) UpdateDocument(c *gin.Context, id string, updated *consulting.Consultation) error {
	args := m.Called(c, id, updated)
	return args.Error(0)
}

func (m *MockDbService) ListDocuments(c *gin.Context) ([]consulting.Consultation, error) {
	return nil, nil
}

func (m *MockDbService) DeleteDocument(c *gin.Context, id string) error {
	return nil
}

func (m *MockDbService) CreateDocument(c *gin.Context, id string, doc *consulting.Consultation) error {
	return nil
}

func TestUpdateConsultation_MissingID(t *testing.T) {
	router := gin.Default()
	handler := consulting.NewConsultationsApi()

	router.PUT("/consultations/:requestId", func(c *gin.Context) {
		c.Set("db_service", &MockDbService{})
		handler.UpdateConsultation(c)
	})

	req, _ := http.NewRequest(http.MethodPut, "/consultations/", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code) // Gin returns 404 if route doesn't match
}
