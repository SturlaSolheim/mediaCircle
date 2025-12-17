package service

import (
	"github.com/SturlaSolheim/mediaCircleBackend/generated"
)

type ListeService interface {
	GetListe() (generated.Liste, error)
}

type ListeServiceImpl struct {
}
func NewListeService(
) ListeService {
	return &ListeServiceImpl{}
}

func (a *ListeServiceImpl) GetListe() (generated.Liste, error) {
	return generated.Liste{}, nil
}

