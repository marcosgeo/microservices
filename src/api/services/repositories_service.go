package services

import (
	"strings"

	"github.com/marcosgeo/microservices/src/api/config"
	"github.com/marcosgeo/microservices/src/api/domain/github"
	"github.com/marcosgeo/microservices/src/api/domain/repositories"
	providers "github.com/marcosgeo/microservices/src/api/providers/github_provider"
	"github.com/marcosgeo/microservices/src/api/utils/errors"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError)
}

var (
	// RepositoryService is an interface for a repository
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

// CreateRepo ...
func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("Invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := providers.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		ID:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}
