package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateResourceTable(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	request, err := http.NewRequest("POST", "/api/v1/table", nil)
	if err != nil {
		t.Errorf("Post heartbeat failed with error %d.", err)
	}

	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("/api/v1/resources failed with error code %d.", response.Code)
	}
}

func TestDropResourceTable(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	request, err := http.NewRequest("DELETE", "/api/v1/table", nil)
	if err != nil {
		t.Errorf("DELETE heartbeat failed with error %d.", err)
	}

	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("/api/v1/resources failed with error code %d.", response.Code)
	}
}

func TestPostItem(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	body := bytes.NewBuffer([]byte(`{"name":"A","enable":false,"state":""}`))
	// body := bytes.NewBuffer([]byte(`{"name":"A","state":{"enable":false,"state":""}}`))

	request, err := http.NewRequest("POST", "/api/v1/resources", body)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("Post heartbeat failed with error %d.", err)
	}

	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	// log.Println(response.Body)

	if response.Code != http.StatusCreated {
		t.Errorf("/api/v1/resources failed with error code %d.", response.Code)
	}

	expectedStringBody := `{"id":1,"name":"A","enable":false,"state":""}`
	// expectedStringBody := `{"id":1,"name":"A","state":{"enable":false,"state":""}}`
	assert.Equal(t, expectedStringBody, response.Body.String())
}

func TestGetItem(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	request, err := http.NewRequest("GET", "/api/v1/resources/1", nil)
	if err != nil {
		t.Errorf("Get heartbeat failed with error %d.", err)
	}

	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	// log.Println(response.Body)

	if response.Code != http.StatusOK {
		t.Errorf("/api/v1/resources failed with error code %d.", response.Code)
	}

	expectedStringBody := `{"id":1,"name":"A","enable":false,"state":""}`
	// expectedStringBody := `{"id":1,"name":"A","state":{"enable":false,"state":""}}`
	assert.Equal(t, expectedStringBody, response.Body.String())
}

func TestGetItems(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	request, err := http.NewRequest("GET", "/api/v1/resources", nil)
	if err != nil {
		t.Errorf("Get heartbeat failed with error %d.", err)
	}

	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	// log.Println(response.Body)

	if response.Code != http.StatusOK {
		t.Errorf("/api/v1/resources failed with error code %d.", response.Code)
	}

	expectedStringBody := `[{"id":1,"name":"A","enable":false,"state":""}]`
	// expectedStringBody := `[{"id":1,"name":"A","state":{"enable":false,"state":""}}]`
	assert.Equal(t, expectedStringBody, response.Body.String())
}

func TestPutItem(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	body := bytes.NewBuffer([]byte(`{"name":"B","enable":false,"state":""}`))
	// body := bytes.NewBuffer([]byte(`"name":"B","state":{"enable":false,"state":""}`))

	request, err := http.NewRequest("PUT", "/api/v1/resources/1", body)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("Put heartbeat failed with error %d.", err)
	}

	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	// log.Println(response.Body)

	if response.Code != http.StatusOK {
		t.Errorf("/api/v1/resources failed with error code %d.", response.Code)
	}

	expectedStringBody := `{"id":1,"name":"B","enable":false,"state":""}`
	// expectedStringBody := `{"id":1,"name":"B","state":{"enable":false,"state":""}}`
	assert.Equal(t, expectedStringBody, response.Body.String())
}

// func TestDeleteItem(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	testRouter := SetupRouter()

// 	request, err := http.NewRequest("DELETE", "/api/v1/resources/1", nil)
// 	request.Header.Set("Content-Type", "application/json")
// 	if err != nil {
// 		t.Errorf("Delete heartbeat failed with error %d.", err)
// 	}

// 	response := httptest.NewRecorder()
// 	testRouter.ServeHTTP(response, request)
// 	// log.Println(response.Body)

// 	if response.Code != http.StatusOK {
// 		t.Errorf("/api/v1/resources failed with error code %d.", response.Code)
// 	}

// 	expectedStringBody := `{"id #1":" deleted"}`
// 	assert.Equal(t, expectedStringBody, response.Body.String())
// }

func TestGetItemHealth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	request, err := http.NewRequest("GET", "/api/v1/resources/1/health", nil)
	if err != nil {
		t.Errorf("Get heartbeat failed with error %d.", err)
	}

	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	// log.Println(response.Body)

	if response.Code != http.StatusOK {
		t.Errorf("/api/v1/resources/ failed with error code %d.", response.Code)
	}

	expectedStringBody := `{"id":1,"name":"B","enable":false,"state":""}`
	// expectedStringBody := `{"id":1,"name":"B","state":{"enable":false,"state":""}}`
	assert.Equal(t, expectedStringBody, response.Body.String())
}

func TestPutItemHealth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	// body := bytes.NewBuffer([]byte("{\"state\": \"Available\", \"name\": \"A\"}"))
	body := bytes.NewBuffer([]byte(`{"name":"A","enable":true,"state":"1"}`))

	request, err := http.NewRequest("PUT", "/api/v1/resources/1/health", body)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("Put heartbeat failed with error %d.", err)
	}

	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	// log.Println(response.Body)

	if response.Code != http.StatusMethodNotAllowed {
		t.Errorf("/api/v1/resources failed with error code %d.", response.Code)
	}
	expectedStringBody := `{"error":"Resource Update is disable"}`

	assert.Equal(t, expectedStringBody, response.Body.String())
}

func TestPutItemEnable(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	body := bytes.NewBuffer([]byte(`{"name":"B","enable":true,"state":""}`))
	// body := bytes.NewBuffer([]byte(`"name":"B","state":{"enable":true,"state":""}`))

	request, err := http.NewRequest("PUT", "/api/v1/resources/1", body)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("Put heartbeat failed with error %d.", err)
	}

	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	// log.Println(response.Body)

	if response.Code != http.StatusOK {
		t.Errorf("/api/v1/resources failed with error code %d.", response.Code)
	}

	expectedStringBody := `{"id":1,"name":"B","enable":true,"state":""}`
	// expectedStringBody := `{"id":1,"name":"B","state":{"enable":true,"state":""}}`
	assert.Equal(t, expectedStringBody, response.Body.String())
}

func TestPutItemHealthEnable(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	body := bytes.NewBuffer([]byte(`{"name":"B","enable":true,"state":"1"}`))
	// body := bytes.NewBuffer([]byte(`{"name":"A","state":{"enable":true,"state":"1"}}`))

	request, err := http.NewRequest("PUT", "/api/v1/resources/1/health", body)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("Put heartbeat failed with error %d.", err)
	}

	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	// log.Println(response.Body)

	if response.Code != http.StatusOK {
		t.Errorf("/api/v1/resources failed with error code %d.", response.Code)
	}
	expectedStringBody := `{"id":1,"name":"B","enable":true,"state":"1"}`
	// expectedStringBody := `{"id":1,"name":"B","state":{"enable":true,"state":"1"}}`

	assert.Equal(t, expectedStringBody, response.Body.String())
}
