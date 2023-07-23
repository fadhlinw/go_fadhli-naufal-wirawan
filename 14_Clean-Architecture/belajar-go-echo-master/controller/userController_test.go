package controller

import (
	"belajar-go-echo/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func TestGetAllUsers(t *testing.T) {
	e := echo.New()

	mockDB := new(MockDB)

	mockDB.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{})

	request := httptest.NewRequest(http.MethodGet, "/users", nil)
	record := httptest.NewRecorder()
	c := e.NewContext(request, record)

	handler := GetAllUsers(mockDB)
	err := handler(c)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, record.Code)

	expectedResponse := echo.Map{
		"data": []model.User{},
	}

	jsonStr, err := json.Marshal(expectedResponse)

	assert.Nil(t, err)
	assert.NotNil(t, jsonStr)
	assert.JSONEq(t, record.Body.String(), string(jsonStr))
}

func TestCreateUser(t *testing.T) {
	e := echo.New()

	mockDb := new(MockDB)

	mockDb.On("Create", mock.Anything, mock.Anything).Return(&gorm.DB{})

	request := httptest.NewRequest(http.MethodPost, "/users", nil)
	record := httptest.NewRecorder()
	c := e.NewContext(request, record)

	handler := CreateUser(mockDb)
	err := handler(c)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, record.Code)

	expectedResponse := echo.Map{
		"data": model.User{
			Email:    "",
			Password: "",
		},
	}

	jsonStr, err := json.Marshal(expectedResponse)

	assert.Nil(t, err)
	assert.NotNil(t, jsonStr)
	assert.JSONEq(t, record.Body.String(), string(jsonStr))
}

func TestBadRequestCreateUser(t *testing.T) {

	e := echo.New()

	mockDB := new(MockDB)

	mockDB.On("Create", mock.Anything).Return(&gorm.DB{})

	requestData := `{
		"email": "
	}`

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(requestData))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	handler := CreateUser(mockDB)
	err := handler(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, res.Code)
}
