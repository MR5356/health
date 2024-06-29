# Go Health Check

[![go-test](https://github.com/MR5356/health/workflows/Go%20Test/badge.svg?query=branch%3Amaster)](https://github.com/MR5356/health/actions?query=branch%3Amaster)
[![go-report](https://goreportcard.com/badge/github.com/MR5356/health)](https://goreportcard.com/report/github.com/MR5356/health)

----

An easy-to-use service status monitoring library in Golang

----

## Installation

```shell
go get github.com/MR5356/health
```

## Example
### URL Check
```go
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
```
Output:
```shell
status: up
rtt: 160ms
result: &{Code:200 Error:<nil>}
```

### Host Check
```go
package main

import (
	"fmt"
	"github.com/MR5356/health/host"
)

func main() {
	// use ping checker
	pingChecker := host.NewPingChecker("baidu.com")
	result := pingChecker.Check()
	fmt.Printf("status: %s\nrtt: %dms\nresult: %+v\n", result.Status, result.GetRTT(), result.GetResult())

	// use ssh checker
	sshChecker := host.NewSSHChecker(&host.HostInfo{
		Host:     "host or ip",
		Port:     22,
		Username: "username",
		Password: "password",
	})
	result = sshChecker.Check()
	fmt.Printf("status: %s\nrtt: %dms\nresult: %+v\n", result.Status, result.GetRTT(), result.GetResult())
}
```
Output:
```shell
status: up
rtt: 9ms
result: &{Addr:39.156.66.10 Error:<nil>}
status: up
rtt: 295ms
result: <nil>
```

### Database Check
```go
package main

import (
	"fmt"
	"github.com/MR5356/health/database"
)

func main() {
	checker := database.NewChecker(database.DBDriverSQLite, "file::memory:?cache=shared")
	result := checker.Check()
	fmt.Printf("status: %s\nrtt: %dms\nresult: %+v\n", result.Status, result.GetRTT(), result.GetResult())
}

```
Output:
```shell
status: up
rtt: 1ms
result: &{Version:3.45.1 Error:<nil>}
```

## Testing
```shell
go test -v ./... -coverprofile=coverage.out
go tool cover -func=coverage.out 
```

## Implementing custom checker
```go
type Checker interface {
	Check() Health
}
```
## Support

If you have questions, reach out to us one way or another.