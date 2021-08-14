package sumolog

import (
	"net/url"
	"strings"

	sumoll "github.com/ushios/sumoll"
)

type SumoLogic struct{
	client *sumoll.HTTPSourceClient
}

func NewSumoLogic(sumoEndpoint *url.URL) (*SumoLogic, error) {
	client, err := sumoll.NewHTTPSourceClient(sumoEndpoint)
	if err != nil {
		return &SumoLogic{}, err
	}

	return &SumoLogic{
		client: client,
	}, nil
}

func (s *SumoLogic) Write(p []byte) (int, error) {
	err := s.client.Send(strings.NewReader(string(p[:])))
	if err != nil {
		// Never fail for now
		return len(p), nil
	}

	return len(p), nil
}