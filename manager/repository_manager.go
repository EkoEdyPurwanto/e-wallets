package manager

import "EEP/e-wallets/repository"

type RepositoryManager interface {
	UsersAccountRepository() repository.UsersAccountRepository
	WalletsRepository() repository.WalletsRepository
}

type repositoryManager struct {
	infraManager InfraManager
}

func NewRepoManager(infraManager InfraManager) RepositoryManager {
	return &repositoryManager{
		infraManager: infraManager,
	}
}

// implement interface here
func (r *repositoryManager) UsersAccountRepository() repository.UsersAccountRepository {
	return repository.NewUsersAccountRepository(r.infraManager.Conn())
}

func (r *repositoryManager) WalletsRepository() repository.WalletsRepository {
	return repository.NewWalletsRepository(r.infraManager.Conn())
}
