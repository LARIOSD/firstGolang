package ports

type TimePort interface {
	GetCurrentTime() (string, error)
}
