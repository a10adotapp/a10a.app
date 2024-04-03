package service

type SomeService struct{}

func NewSomeService() SomeService {
	return SomeService{}
}

func (s SomeService) CreateRecord() error {
	return nil
}
