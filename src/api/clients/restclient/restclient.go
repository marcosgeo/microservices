package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enableMocks = false
	mocks       = make(map[string]*Mock)
)

type Mock struct {
	Url        string
	HttpMethod string
	Response   *http.Response
	Err        error
}

func getMockID(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

// StartMockups ...
func StartMockups() {
	enableMocks = true
}

// StopMockups ...
func StopMockups() {
	enableMocks = false
}

// FlushMockups ...
func FlushMockups() {
	mocks = make(map[string]*Mock)
}

// AddMockup ...
func AddMockup(mock Mock) {
	// sets the key:value of the mock map
	mocks[getMockID(mock.HttpMethod, mock.Url)] = &mock
	fmt.Println(mocks)
}

// Post makes a http POST request
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {

	if enableMocks {
		mock := mocks[getMockID(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("No mockup found for the given request")
		}
		return mock.Response, mock.Err
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}
