package service

import "github.com/SturlaSolheim/mediaCircleBackend/generated"

type AnmeldelseService interface {
	GetAlleAnmeldelserForEnBruker() ([]generated.AnmeldelseSlim, error)
}

type AnmeldelseServiceImpl struct {
}
func NewAnmeldelseService(
) AnmeldelseService {
	return &AnmeldelseServiceImpl{}
}

func (a *AnmeldelseServiceImpl) GetAlleAnmeldelserForEnBruker() ([]generated.AnmeldelseSlim, error) {
	return make([]generated.AnmeldelseSlim, 1), nil;
}
