package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/marcosgeo/microservices/src/api/clients/restclient"
	"github.com/marcosgeo/microservices/src/api/domain/repositories"
	"github.com/marcosgeo/microservices/src/api/utils/errors"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepo_InvalidInputName(t *testing.T) {
	request := repositories.CreateRepoRequest{} // creates an empty request

	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "Invalid repository name", err.Message())
}

func TestCreateRepo_ErrorFromGithub(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://developer.github.com/docs"}`)),
		},
	})

	request := repositories.CreateRepoRequest{Name: "testing"}

	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Requires authentication", err.Message())
}

func TestCreateRepo_NoError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing", "owner": {"login": "marcosgeo"}}`)),
		},
	})

	request := repositories.CreateRepoRequest{Name: "testing"}

	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 123, result.ID)
	assert.EqualValues(t, "testing", result.Name)
	assert.EqualValues(t, "marcosgeo", result.Owner)
}

func TestCreateRepoConcurrent_InvalidRequest(t *testing.T) {
	request := repositories.CreateRepoRequest{}

	output := make(chan repositories.CreateRepositoriesResult)

	service := repoService{}
	go service.createRepoConcurrent(request, output)

	result := <-output

	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
	assert.EqualValues(t, "Invalid repository name", result.Error.Message())
}

func TestCreateRepoConcurrent_ErrorFromGithub(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"Message": "Requires authentication", "documentation_url": "https://developer.github.com/doc"}`)),
		},
	})

	request := repositories.CreateRepoRequest{Name: "testing"}

	output := make(chan repositories.CreateRepositoriesResult)

	service := repoService{}
	go service.createRepoConcurrent(request, output)

	result := <-output

	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusUnauthorized, result.Error.Status())
	assert.EqualValues(t, "Requires authentication", result.Error.Message())
}

func TestCreateRepoConcurrent_NoError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing", "owner": {"login": "marcosgeo"}}`)),
		},
	})

	request := repositories.CreateRepoRequest{Name: "testing"}

	output := make(chan repositories.CreateRepositoriesResult)

	service := repoService{}
	go service.createRepoConcurrent(request, output)

	result := <-output

	assert.NotNil(t, result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Response)
	assert.EqualValues(t, 123, result.Response.ID)
	assert.EqualValues(t, "testing", result.Response.Name)
	assert.EqualValues(t, "marcosgeo", result.Response.Owner)
}

func TestHundeResults_Concurrent(t *testing.T) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)

	var wg sync.WaitGroup

	service := repoService{}
	go service.handleRepoResults(&wg, input, output)

	wg.Add(1)
	go func() {
		input <- repositories.CreateRepositoriesResult{
			Error: errors.NewBadRequestError("Invalid repository name"),
		}
	}()

	wg.Wait()
	close(input)

	result := <-output

	assert.NotNil(t, result)
	assert.EqualValues(t, 0, result.StatusCode)

	assert.EqualValues(t, 1, len(result.Results))
	assert.NotNil(t, result.Results[0].Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())
	assert.EqualValues(t, "Invalid repository name", result.Results[0].Error.Message())
}

func TestCreateRepos_InvalidRequests(t *testing.T) {
	// inicialization
	requests := []repositories.CreateRepoRequest{
		{},
		{Name: " "},
	}

	// execution
	result, err := RepositoryService.CreateRepos(requests)

	// validation
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 2, len(result.Results))
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())
	assert.EqualValues(t, "Invalid repository name", result.Results[0].Error.Message())
	assert.EqualValues(t, http.StatusBadRequest, result.Results[1].Error.Status())
	assert.EqualValues(t, "Invalid repository name", result.Results[1].Error.Message())
}

func TestCreateRepos_OneSuccessOneFail(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing", "owner": {"login":"marcosgeo"}}`)),
		},
	})

	requests := []repositories.CreateRepoRequest{
		{Name: "testing0"},
		{Name: "testing1"},
	}

	result, err := RepositoryService.CreateRepos(requests)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusPartialContent, result.StatusCode)
	assert.EqualValues(t, 2, len(result.Results))

	assert.Nil(t, result.Results[0].Error)
	assert.EqualValues(t, 123, result.Results[0].Response.ID)
	assert.EqualValues(t, "testing", result.Results[0].Response.Name)
	assert.EqualValues(t, "marcosgeo", result.Results[0].Response.Owner)

	assert.Nil(t, result.Results[1].Error)
	// assert.EqualValues(t, 123, result.Results[1].Response.ID)
	// assert.EqualValues(t, "testing", result.Results[1].Response.Name)
	// assert.EqualValues(t, "marcosgeo", result.Results[1].Response.Owner)
}
