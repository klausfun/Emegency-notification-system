package service

type Sending interface {
	SendEmail() error
}

type Service struct {
	Sending
}

func NewService() *Service {
	return &Service{
		Sending: NewSendingService(),
	}
}
