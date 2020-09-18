package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthRoute(t *testing.T) {
	testServer := httptest.NewServer(SetupServer())
	defer testServer.Close()

	resp, err := http.Get(fmt.Sprintf("%s/health", testServer.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected content-type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}

}

func TestRegistrationRoute(t *testing.T) {
	requestBody, err := json.Marshal(map[string]string{
		"username": "test00",
		"email":    "test.user@test.com",
		"password": "t3stp4ssw0rd",
	})

	if err != nil {
		t.Fatalf("Expected no error while parsing json, got %v", err)
	}

	testServer := httptest.NewServer(SetupServer())
	defer testServer.Close()

	resp, err := http.Post(fmt.Sprintf("%s/user/create", testServer.URL), "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected content-type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}

	var jsonResponse map[string]int
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	if err != nil {
		t.Fatalf("Unable to decode json received")
	}

	_, ok = jsonResponse["id"]
	if !ok {
		t.Fatalf("Expected to recieve \"ID\" in the JSON")
	}
}

func TestAuthenticationRoute(t *testing.T) {
	requestBody, err := json.Marshal(map[string]string{
		"username": "test00",
		"password": "t3stp4ssw0rd",
	})

	if err != nil {
		t.Fatalf("Expected no error while parsing json, got %v", err)
	}

	testServer := httptest.NewServer(SetupServer())
	defer testServer.Close()

	resp, err := http.Post(fmt.Sprintf("%s/user/authenticate", testServer.URL), "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected content-type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}

	var jsonResponse map[string]string
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	if err != nil {
		t.Fatalf("Unable to decode json received")
	}

	_, ok = jsonResponse["accessToken"]
	if !ok {
		t.Fatalf("Expected to recieve \"accessToken\" in the JSON")
	}
}

func ExchangeRateCreationRoute(t *testing.T) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"currency": "EUR",
		"rate":     312341,
	})

	if err != nil {
		t.Fatalf("Expected no error while parsing json, got %v", err)
	}

	testServer := httptest.NewServer(SetupServer())
	defer testServer.Close()

	resp, err := http.Post(fmt.Sprintf("%s/exhangerate/create", testServer.URL), "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected content-type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}

	var jsonResponse map[string]int
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	if err != nil {
		t.Fatalf("Unable to decode json received")
	}

	_, ok = jsonResponse["id"]
	if !ok {
		t.Fatalf("Expected to recieve \"id\" in the JSON")
	}
}

func CategoryCreationTest(t *testing.T) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"name": "test category",
	})

	if err != nil {
		t.Fatalf("Expected no error while parsing json, got %v", err)
	}

	testServer := httptest.NewServer(SetupServer())
	defer testServer.Close()

	resp, err := http.Post(fmt.Sprintf("%s/category/create", testServer.URL), "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected content-type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}

	var jsonResponse map[string]int
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	if err != nil {
		t.Fatalf("Unable to decode json received")
	}

	_, ok = jsonResponse["id"]
	if !ok {
		t.Fatalf("Expected to recieve \"id\" in the JSON")
	}
}
