package host

import (
	ping "github.com/MR5356/go-ping"
	"github.com/MR5356/health"
	"time"
)

type PingChecker struct {
	host    string
	timeout time.Duration
}

type PingResult struct {
	Addr  string `json:"addr"`
	Error error  `json:"error"`
}

func NewPingChecker(host string) *PingChecker {
	return &PingChecker{
		host:    host,
		timeout: time.Second * 5,
	}
}

func NewPingCheckerWithTimeout(host string, timeout time.Duration) *PingChecker {
	return &PingChecker{
		host:    host,
		timeout: timeout,
	}
}

func (pc *PingChecker) Check() (result *health.Health) {
	result = health.NewHealth()
	pinger, err := ping.NewPinger(pc.host)
	if err != nil {
		result.Down()
		result.SetResult(&PingResult{Error: err})
		return
	}

	pinger.Count = 1
	pinger.Timeout = pc.timeout

	if err = pinger.Run(); err != nil {
		result.Down()
		result.SetResult(&PingResult{Error: err})
	} else {
		stats := pinger.Statistics()
		result.SetRTT(stats.AvgRtt.Milliseconds())
		result.Up()
		result.SetResult(&PingResult{Addr: stats.IPAddr.String()})
	}
	return
}
