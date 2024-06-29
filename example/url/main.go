package main

import (
	"fmt"
	"github.com/MR5356/health/url"
)

func main() {
	checker := url.NewChecker("https://toodo.fun")

	result := checker.Check()
	fmt.Printf("status: %s\nrtt: %dms\nresult: %+v\n", result.Status, result.GetRTT(), result.GetResult())
}
