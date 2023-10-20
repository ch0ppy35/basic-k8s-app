package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response MessageResponse
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Hello World!", response.Message)
}

func TestMessageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/message", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(messageHandler)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response MessageResponse
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Yay, K8s!", response.Message)
}

func TestPodNameHandler(t *testing.T) {
	// Set the MY_POD_NAME environment variable for the test
	os.Setenv("MY_POD_NAME", "TestPod")

	req, err := http.NewRequest("GET", "/pod_name", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(podNameHandler)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response PodNameResponse
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "TestPod", response.PodName)

	// Clean up the environment variable after the test
	os.Unsetenv("MY_POD_NAME")
}

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthz", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(healthCheckHandler)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response HealthCheckResponse
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, true, response.Status)
}
