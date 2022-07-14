package service

type ServiceConfig struct {
	DictStorage DictionaryStorage
}

type Service struct {
	Dict *DictionaryService
}

func NewService(cfg *ServiceConfig) *Service {
	return &Service{
		Dict: NewDictionaryService(cfg.DictStorage),
	}
}
