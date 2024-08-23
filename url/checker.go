package url

import (
	"github.com/MR5356/health"
	"net/http"
	"strings"
	"time"
)

type Checker struct {
	url     string
	timeout time.Duration
}

type Result struct {
	Code  int   `json:"code"`
	Error error `json:"error"`
}

func NewChecker(url string) *Checker {
	return NewCheckerWithTimeout(url, time.Second*5)
}

func NewCheckerWithTimeout(url string, timeout time.Duration) *Checker {
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	return &Checker{
		url:     url,
		timeout: timeout,
	}
}

func (c *Checker) Check() (result *health.Health) {
	result = health.NewHealth()

	client := http.Client{
		Timeout: c.timeout,
	}

	startT := time.Now()
	resp, err := client.Head(c.url)
	rtt := time.Since(startT).Milliseconds()
	result.SetRTT(rtt)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		result.Down()
		result.SetResult(&Result{
			Code:  0,
			Error: err,
		})
		return
	}
	if resp != nil {
		result.SetResult(&Result{Code: resp.StatusCode})
		if resp.StatusCode == http.StatusOK {
			result.Up()
		} else {
			result.Down()
		}
	}
	return result
}
