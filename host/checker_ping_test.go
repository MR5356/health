package host

import (
	"github.com/MR5356/health"
	"testing"
	"time"
)

func TestNewPingChecker(t *testing.T) {
	host := "toodo.fun"

	pingChecker := NewPingChecker(host)
	if pingChecker.host != host {
		t.Errorf("Expected host %s, got %s", host, pingChecker.host)
	}
	if pingChecker.timeout != time.Second*5 {
		t.Errorf("Expected timeout %v, got %v", time.Second*5, pingChecker.timeout)
	}
}

func TestNewPingCheckerWithTimeout(t *testing.T) {
	url := "toodo.fun"
	timeout := time.Second * 10

	pingChecker := NewPingCheckerWithTimeout(url, timeout)
	if pingChecker.host != url {
		t.Errorf("Expected host %s, got %s", url, pingChecker.host)
	}
	if pingChecker.timeout != timeout {
		t.Errorf("Expected timeout %v, got %v", timeout, pingChecker.timeout)
	}
}

func TestPingChecker_Check_Success(t *testing.T) {
	host := "baidu.com"

	pingChecker := NewPingChecker(host)

	result := pingChecker.Check()
	if result.Status != health.StatusUp {
		t.Errorf("Expected status %s, got %s", health.StatusDown, result.Status)
	}
	if !result.HasResult() {
		t.Errorf("Expected to have result, but got %v", result.GetResult())
	}
}

func TestPingChecker_Check_Failure(t *testing.T) {
	host := "toodo.fun:2345"

	pingChecker := NewPingChecker(host)

	result := pingChecker.Check()
	if result.Status != health.StatusDown {
		t.Errorf("Expected status %s, got %s", health.StatusDown, result.Status)
	}
	if !result.HasResult() {
		t.Errorf("Expected to have result, but got %v", result.GetResult())
	}
}

func TestPingChecker_Check_Empty(t *testing.T) {
	host := "toodo.fun"

	pingChecker := NewPingCheckerWithTimeout(host, 1)

	result := pingChecker.Check()
	if result.Status != health.StatusDown {
		t.Errorf("Expected status %s, got %s", health.StatusDown, result.Status)
	}
	if !result.HasResult() {
		t.Errorf("Expected to have result, but got %v", result.GetResult())
	}
}
