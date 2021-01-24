package base

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

// PutHttpRequestInput configures a PUT request.
type PutHttpRequestInput struct {
	Body             []byte
	ValidStatusCodes []int
	ValidStatusFunc  ValidStatusFunc
	Uri              Uri
}

// GetValidStatusCodes returns a []int of status codes considered valid for a PUT request.
func (i PutHttpRequestInput) GetValidStatusCodes() []int {
	return i.ValidStatusCodes
}

// GetValidStatusFunc returns a function used to evaluate whether the response to a PUT request is considered valid.
func (i PutHttpRequestInput) GetValidStatusFunc() ValidStatusFunc {
	return i.ValidStatusFunc
}

// Put performs a PUT request.
func (c Client) Put(ctx context.Context, input PutHttpRequestInput) (*http.Response, int, error) {
	var status int
	url, err := c.buildUri(input.Uri)
	if err != nil {
		return nil, status, fmt.Errorf("unable to make request: %v", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(input.Body))
	if err != nil {
		return nil, status, err
	}
	resp, status, _, err := c.performRequest(req, input)
	if err != nil {
		return nil, status, err
	}
	return resp, status, nil
}
