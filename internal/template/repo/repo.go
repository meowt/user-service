package repo

import (
	"project/internal/storage"
)

type TemplateRepoImpl struct {
	Storage *storage.Storage
}

func SetupRepo(s *storage.Storage) *TemplateRepoImpl {
	return &TemplateRepoImpl{
		Storage: s,
	}
}

func (r *TemplateRepoImpl) Example() (err error) {
	return nil
}
