package bar

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetBar(id string) *Bar {
	return &Bar{
		Id: id,
	}
}

type Bar struct {
	Id string `json:"id"`
}
