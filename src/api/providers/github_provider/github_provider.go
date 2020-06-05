package providers

import (
	"fmt"

	github "github.com/marcosgeo/microservices/src/api/domain/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

// CreateRepo creates a github repository
func CreateRepo(accessToken string, request github.CreateRepoRequest) (github.CreateRepoResponse, github.GithubErrorResponse) {
	// Autorization: token 98asçldkj9321u4lçkasdjtoqwe4
	header := getAuthorizationHeader(accessToken)

	fmt.Println(header)
	return github.CreateRepoResponse{}, github.GithubErrorResponse{}
}
