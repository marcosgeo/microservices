package services

import (
	"net/http"
	"sync"

	"github.com/marcosgeo/microservices/src/api/config"
	"github.com/marcosgeo/microservices/src/api/domain/github"
	"github.com/marcosgeo/microservices/src/api/domain/repositories"
	providers "github.com/marcosgeo/microservices/src/api/providers/github_provider"
	"github.com/marcosgeo/microservices/src/api/utils/errors"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError)
	CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.APIError)
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
	if err := input.Validate(); err != nil {
		return nil, err
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

// CreateRepos ...
func (s *repoService) CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.APIError) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)

	var wg sync.WaitGroup
	go s.handleRepoResults(&wg, input, output)

	for _, current := range requests {
		wg.Add(1)
		go s.createRepoConcurrent(current, input)
	}

	wg.Wait() // wait group, waits until all requests are processed
	close(input)

	result := <-output

	successCreations := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successCreations++
		}
	}

	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreations == len(requests) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}

	return result, nil
}

func (s *repoService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse

	for repositoryResult := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: repositoryResult.Response,
			Error:    repositoryResult.Error,
		}
		results.Results = append(results.Results, repoResult)

		wg.Done()
	}
	output <- results
}

func (s *repoService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}
	result, err := s.CreateRepo(input)
	if err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}
	output <- repositories.CreateRepositoriesResult{Response: result}
}
