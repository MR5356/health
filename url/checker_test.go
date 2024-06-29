package url

import (
	"github.com/MR5356/health"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewChecker(t *testing.T) {
	url := "http://example.com"
	checker := NewChecker(url)
	if checker.url != url {
		t.Errorf("Expected URL %s, got %s", url, checker.url)
	}
	if checker.timeout != time.Second*5 {
		t.Errorf("Expected timeout %v, got %v", time.Second*5, checker.timeout)
	}
}

func TestNewCheckerWithTimeout(t *testing.T) {
	url := "http://example.com"
	timeout := time.Second * 10
	checker := NewCheckerWithTimeout(url, timeout)
	if checker.url != url {
		t.Errorf("Expected URL %s, got %s", url, checker.url)
	}
	if checker.timeout != timeout {
		t.Errorf("Expected timeout %v, got %v", timeout, checker.timeout)
	}
}

func TestChecker_Check_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	checker := NewChecker(server.URL)
	result := checker.Check()

	if result.Status != health.StatusUp {
		t.Errorf("Expected status %s, got %s", health.StatusUp, result.Status)
	}
	if !result.HasResult() {
		t.Errorf("Expected to have result, but got none")
	}

	res, ok := result.GetResult().(*Result)
	if !ok {
		t.Fatalf("Expected result type *Result, got %T", result.GetResult())
	}
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}
}

func TestChecker_Check_Failure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	checker := NewChecker(server.URL)
	result := checker.Check()

	if result.Status != health.StatusDown {
		t.Errorf("Expected status %s, got %s", health.StatusDown, result.Status)
	}
	if !result.HasResult() {
		t.Errorf("Expected to have result, but got none")
	}

	res, ok := result.GetResult().(*Result)
	if !ok {
		t.Fatalf("Expected result type *Result, got %T", result.GetResult())
	}
	if res.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, res.Code)
	}
}

func TestChecker_Check_BadGateWay(t *testing.T) {
	t.Parallel()
	checker := NewChecker("http://invalid-url")
	result := checker.Check()

	if result.Status != health.StatusDown {
		t.Errorf("Expected status %s, got %s", health.StatusDown, result.Status)
	}
	if !result.HasResult() {
		t.Errorf("Expected to have result, but got none")
	}

	res, ok := result.GetResult().(*Result)
	if !ok {
		t.Fatalf("Expected result type *Result, got %T", result.GetResult())
	}
	if res.Code != 0 && res.Code != 502 {
		t.Errorf("Expected status code 0, got %d", res.Code)
	}
}

func TestChecker_Check_Error(t *testing.T) {
	t.Parallel()
	checker := NewChecker("invalid-url")
	result := checker.Check()

	if result.Status != health.StatusDown {
		t.Errorf("Expected status %s, got %s", health.StatusDown, result.Status)
	}
	if !result.HasResult() {
		t.Errorf("Expected to have result, but got none")
	}

	res, ok := result.GetResult().(*Result)
	if !ok {
		t.Fatalf("Expected result type *Result, got %T", result.GetResult())
	}
	if res.Code != 0 {
		t.Errorf("Expected status code 0, got %d", res.Code)
	}
	if res.Error == nil {
		t.Errorf("Expected error, got nil")
	}
}
