package usecase

type TemplateRepoService interface {
}

type TemplateUsecaseImpl struct {
	Repo TemplateRepoService
}

func SetupUsecase(r TemplateRepoService) *TemplateUsecaseImpl {
	return &TemplateUsecaseImpl{
		Repo: r,
	}
}

func (u *TemplateUsecaseImpl) Example() (err error) {
	return nil
}
