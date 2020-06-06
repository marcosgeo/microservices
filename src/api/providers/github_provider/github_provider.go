package providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/marcosgeo/microservices/src/api/clients/restclient"
	github "github.com/marcosgeo/microservices/src/api/domain/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

// CreateRepo creates a github repository
func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	// Autorization: token 98asçldkj9321u4lçkasdjtoqwe4
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	response, err := restclient.Post(urlCreateRepo, request, headers)
	fmt.Println(response)
	fmt.Println(err)
	if err != nil {
		log.Println(fmt.Sprintf("Error when trying to create new repo in ginhub: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error()}

	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid response body",
		}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Invalid json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse

	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("Error when trying to unmarshal create respo ssuccessful response: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error when trying to unmarshal github.CreateRespoResponse",
		}
	}
	return &result, nil
}
