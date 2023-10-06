package manager

type RepositoryManager interface {
}

type repositoryManager struct {
	infraManager InfraManager
}

func NewRepoManager(infraManager InfraManager) RepositoryManager {
	return &repositoryManager{
		infraManager: infraManager,
	}
}

//implement interface here
