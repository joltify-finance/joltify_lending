package constants

import "fmt"

const (
	UnexpectedResponseStatusMessage = "Unexpected response status code of:"
)

var ErrRateLimiting = fmt.Errorf("status 429 - rate limit exceeded")
