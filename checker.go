package health

type Checker interface {
	Check() *Health
}
