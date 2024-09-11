package domain

type TimeService interface {
	GetCurrentTime() (string, error)
}
