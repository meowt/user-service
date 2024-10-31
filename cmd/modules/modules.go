package modules

import (
	"project/internal/api"
	"project/internal/storage"
	templateApi "project/internal/template/api"
	templateRepo "project/internal/template/repo"
	templateUsecase "project/internal/template/usecase"
)

func Setup(s *storage.Storage) *ApiModule {
	repo := SetupRepoModule(s)
	uc := SetupUsecaseModule(repo)
	api := SetupApiModule(uc)
	return api
}

type RepoModule struct {
	template *templateRepo.TemplateRepoImpl
}

func SetupRepoModule(s *storage.Storage) *RepoModule {
	return &RepoModule{
		template: templateRepo.SetupRepo(s),
	}
}

type UsecaseModule struct {
	template *templateUsecase.TemplateUsecaseImpl
}

func SetupUsecaseModule(r *RepoModule) *UsecaseModule {
	return &UsecaseModule{
		template: templateUsecase.SetupUsecase(r),
	}
}

type ApiModule struct {
	template *templateApi.TemplateApiImpl
}

func SetupApiModule(uc *UsecaseModule) *ApiModule {
	return &ApiModule{
		template: templateApi.SetupUsecase(uc),
	}
}

func (a *ApiModule) InitRoutes(r *api.Router) {
	a.template.InitRoutes(r)
}
