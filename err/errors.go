package err

import "errors"

var (
	FailedToParseURI        = errors.New("failed to parse template")
	FailedToMakeHTTPRequest = errors.New("failed to make http request")
	FailedToParseJson       = errors.New("failed to parse JSON")
)
