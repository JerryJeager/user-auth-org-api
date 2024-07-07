package http

import (
	"context"
	"encoding/json"
	// "errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JerryJeager/user-auth-org-api/internal/service/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock the OrgSv interface
type MockOrgSv struct {
	mock.Mock
}
type testRes struct {
	status  string
	message string
	data    models.OrganisationsRes
}

func (m *MockOrgSv) CreateOrganisation(ctx context.Context, org *models.Organisation, userID uuid.UUID) (*models.Organisation, error) {
	args := m.Called(ctx, org, userID)
	return args.Get(0).(*models.Organisation), args.Error(1)
}

func (m *MockOrgSv) CreateOrgMember(ctx context.Context, orgID, userID uuid.UUID) error {
	args := m.Called(ctx, orgID, userID)
	return args.Error(0)
}

func (m *MockOrgSv) GetOrganisation(ctx context.Context, orgID uuid.UUID) (*models.Organisation, error) {
	args := m.Called(ctx, orgID)
	return args.Get(0).(*models.Organisation), args.Error(1)
}

func (m *MockOrgSv) GetOrganisations(ctx context.Context, userID uuid.UUID) (*models.Organisations, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(*models.Organisations), args.Error(1)
}

func TestGetOrganisations(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockOrgSv)
	controller := NewOrgController(mockService)

	userID := uuid.New()
	orgs := make([]struct {
		status  string
		message string
		data    models.OrganisationsRes
	}, 1)
	orgs[0] = testRes{
		status:  "success",
		message: "get all organisations successful",
		data: models.OrganisationsRes{
			Organisation: models.Organisations{},
		},
	}
	mockService.On("GetOrganisations", mock.Anything, userID).Return(orgs, nil)

	req, err := http.NewRequest(http.MethodGet, "/api/organisations", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("user_id", userID.String())

	controller.GetOrganisations(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response["status"])
	assert.Equal(t, "get all organisations successful", response["message"])

	data := response["data"].(map[string]interface{})
	orgList := data["Organisation"].([]interface{})
	assert.Len(t, orgList, 1)
	org := orgList[0].(map[string]interface{})
	assert.Equal(t, "Test Org 1", org["name"])
	assert.Equal(t, "Description 1", org["description"])
}
