package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldShouldPass(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()
	testClient := testServer.Client()

	fmt.Println(testServer.URL)
	testClient.Get(testServer.URL)

	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "HELLO WORLD!", string(body))
	assert.Equal(t, 200, response.StatusCode)
}
func TestHelloWorldShouldFail(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()
	testClient := testServer.Client()

	body := strings.NewReader("some body")

	response, err := testClient.Post(testServer.URL, "applicaition/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, response.StatusCode)
}
func TestHealthShouldPass(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()
	testClient := testServer.Client()
	
	fmt.Println(testServer.URL)
	testClient.Get(testServer.URL)
	
	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "OK", string(body))
	assert.Equal(t, 200, response.StatusCode)
}
func TestHealthShouldFail(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()
	testClient := testServer.Client()

	body := strings.NewReader("some body")

	response, err := testClient.Post(testServer.URL, "applicaition/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, response.StatusCode)
}
func TestNewEndpointShouldPass(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleNewEndpoint))
	defer testServer.Close()
	testClient := testServer.Client()

	fmt.Println(testServer.URL)
	testClient.Get(testServer.URL)

	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "HELLO WORLD!", string(body))
	assert.Equal(t, 200, response.StatusCode)
}
func TestNewEndpointShouldFail(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()
	testClient := testServer.Client()

	body := strings.NewReader("some body")

	response, err := testClient.Post(testServer.URL, "applicaition/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, response.StatusCode)
}
