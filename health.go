package health

import (
	"github.com/MR5356/health/utils"
)

const (
	StatusUp      = Status("up")
	StatusDown    = Status("down")
	StatusUnknown = Status("unknown")
)

type Status string

type Health struct {
	Status Status `json:"status"`
	RTT    int64  `json:"rtt"`
	Result any    `json:"result"`
}

func NewHealth() *Health {
	return &Health{
		Status: StatusUnknown,
		RTT:    0,
		Result: nil,
	}
}

func (h *Health) Up() *Health {
	h.Status = StatusUp
	return h
}

func (h *Health) Down() *Health {
	h.Status = StatusDown
	return h
}

func (h *Health) Unknown() *Health {
	h.Status = StatusUnknown
	return h
}

func (h *Health) SetResult(result any) *Health {
	h.Result = result
	return h
}

func (h *Health) SetRTT(rtt int64) *Health {
	if rtt < 0 {
		rtt = 0.0
	}
	h.RTT = rtt
	return h
}

func (h *Health) IsUp() bool {
	return h.Status == StatusUp
}

func (h *Health) IsDown() bool {
	return h.Status == StatusDown
}

func (h *Health) IsUnknown() bool {
	return h.Status == StatusUnknown
}

func (h *Health) HasResult() bool {
	return !utils.IsZeroValue(h.Result)
}

func (h *Health) HasRTT() bool {
	return h.RTT > 0
}

func (h *Health) GetResult() any {
	return h.Result
}

func (h *Health) GetRTT() int64 {
	return h.RTT
}
