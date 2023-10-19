package manager

import "EEP/e-wallets/usecase"

type UseCaseManager interface {
	UsersAccountUC() usecase.UsersAccountUseCase
	WalletsUC() usecase.WalletsUseCase
}

type useCaseManager struct {
	repositoryManager RepositoryManager
}

func NewUseCaseManager(repositoryManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repositoryManager: repositoryManager,
	}
}

// implement interface here
func (u *useCaseManager) UsersAccountUC() usecase.UsersAccountUseCase {
	return usecase.NewUsersAccountUseCase(u.repositoryManager.UsersAccountRepository(), u.WalletsUC())
}

func (u *useCaseManager) WalletsUC() usecase.WalletsUseCase {
	return usecase.NewWalletsUseCase(u.repositoryManager.WalletsRepository())
}
