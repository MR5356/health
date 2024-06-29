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
