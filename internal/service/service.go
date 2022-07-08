package service

type ServiceConfig struct {
	DictStorage DictionaryStorage
}

type Service struct {
	DictService *DictionaryService
}

func NewService(cfg *ServiceConfig) *Service {
	return &Service{
		DictService: NewDictionaryService(cfg.DictStorage),
	}
}