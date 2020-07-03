package repositories

import (
	"strings"

	"github.com/marcosgeo/microservices/src/api/utils/errors"
)

// CreateRepoRequest ...
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Validate ..
func (r *CreateRepoRequest) Validate() errors.APIError {
	r.Name = strings.TrimSpace(r.Name)
	if r.Name == "" {
		return errors.NewBadRequestError("Invalid repository name")
	}
	return nil
}

// CreateRepoResponse ...
type CreateRepoResponse struct {
	ID    int64  `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

type CreateReposResponse struct {
	StatusCode int                        `json:"status"`
	Results    []CreateRepositoriesResult `json:"result"`
}

type CreateRepositoriesResult struct {
	Response *CreateRepoResponse `json:"repo"`
	Error    errors.APIError     `json:"error"`
}
