package service

type SendingService struct {
}

func NewSendingService() *SendingService {
	return &SendingService{}
}

func (s *SendingService) SendEmail() error {
	return nil
}
