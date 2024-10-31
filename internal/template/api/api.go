package api

import (
	"net/http"

	"project/internal/api"
)

type TemplateUsecaseService interface {
}

type TemplateApiImpl struct {
	Usecase TemplateUsecaseService
}

func SetupUsecase(u TemplateUsecaseService) *TemplateApiImpl {
	return &TemplateApiImpl{
		Usecase: u,
	}
}

func (a *TemplateApiImpl) InitRoutes(router *api.Router) {
}

func (a *TemplateApiImpl) Example(w http.ResponseWriter, r *http.Request) {
}
