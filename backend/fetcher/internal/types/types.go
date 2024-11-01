package types

import (
	"github.com/google/go-github/v66/github"
	"github.com/gregjones/httpcache"
	"net/http"
)

type TransportWithModifiedSince struct {
	ModifiedSince string
}

func (t *TransportWithModifiedSince) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	if t.ModifiedSince != "" {
		req.Header.Set("If-Modified-Since", t.ModifiedSince)
	}

	return http.DefaultTransport.RoundTrip(req)
}

func NewGithubClientWithModifiedSince(modifiedSince string, token string) (client *github.Client) {
	transportWithModifiedSince := &TransportWithModifiedSince{
		ModifiedSince: modifiedSince,
	}

	transport := &httpcache.Transport{
		Transport:           transportWithModifiedSince,
		Cache:               httpcache.NewMemoryCache(),
		MarkCachedResponses: true,
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	client = github.NewClient(httpClient).WithAuthToken(token)

	return
}
