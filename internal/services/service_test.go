// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"quotes/internal/model"
	"strings"
	"testing"
)

func TestFallback(t *testing.T) {
	expected := "Unfortunately, our APIs aren't available right now. Try again later!\n- Every single app in existence"
	if Fallback() != expected {
		t.Errorf("Fallback() = %q, want %q", Fallback(), expected)
	}
}

func TestSaveQuoteFile(t *testing.T) {
	quote := "Test quote\n- Test Author"
	err := SaveQuoteFile(quote)
	if err != nil {
		t.Fatalf("SaveQuoteFile() error = %v", err)
	}

	found := false
	for i := 1; i <= 10; i++ {
		filename := "quote_" + fmt.Sprint(i) + ".txt"
		if _, err := os.Stat(filename); err == nil {
			data, _ := os.ReadFile(filename)
			if strings.Contains(string(data), quote) {
				found = true
				os.Remove(filename)
				break
			}
		}
	}
	if !found {
		t.Error("Quote file not found or content mismatch.")
	}
}

func TestContactZenquotesAPIWithClient(t *testing.T) {
	mockResp := []model.ZenquoteResponse{{Content: "Mock Zenquote", Author: "Mock Author"}}
	respBytes, _ := json.Marshal(mockResp)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(respBytes)
	}))
	defer server.Close()

	origURL := model.ZENQUOTES_URL
	model.ZENQUOTES_URL = server.URL
	defer func() { model.ZENQUOTES_URL = origURL }()

	client := server.Client()
	result := ContactZenquotesAPIWithClient(client)
	if !strings.Contains(result, "Mock Zenquote") || !strings.Contains(result, "Mock Author") {
		t.Errorf("ContactZenquotesAPIWithClient() = %q, want quote and author", result)
	}
}

func TestContactQuotableAPIWithClient(t *testing.T) {
	mockResp := model.QuotableResponse{Content: "Mock Quotable", Author: "Quotable Author"}
	respBytes, _ := json.Marshal(mockResp)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(respBytes)
	}))
	defer server.Close()

	origURL := model.QUOTABLE_URL
	model.QUOTABLE_URL = server.URL
	defer func() { model.QUOTABLE_URL = origURL }()

	client := server.Client()
	result := ContactQuotableAPIWithClient(client)
	if !strings.Contains(result, "Mock Quotable") || !strings.Contains(result, "Quotable Author") {
		t.Errorf("ContactQuotableAPIWithClient() = %q, want quote and author", result)
	}
}
