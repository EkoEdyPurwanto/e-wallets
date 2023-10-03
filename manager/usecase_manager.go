package manager

type UseCaseManager interface {
}

type useCaseManager struct {
	repositoryManager RepositoryManager
}

func NewUseCaseManager(repositoryManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repositoryManager: repositoryManager,
	}
}

//implement interface here
